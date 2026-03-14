package query

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jszwec/csvutil"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/encrypt"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/objectstore"
)

const (
	batchSize    int32  = 100
	exportBucket string = "exports"
)

type Query struct {
	storage   Storage
	encryptor encrypt.Encryptor
	uploader  objectstore.Uploader
}

func NewQuery(
	storage Storage,
	encryptor encrypt.Encryptor,
	uploader objectstore.Uploader,
) *Query {
	return &Query{
		storage:   storage,
		encryptor: encryptor,
		uploader:  uploader,
	}
}

func (q *Query) Get(
	ctx context.Context,
	id uuid.UUID,
) (*RecipeWithUser, error) {
	result, err := q.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (q *Query) Export(
	ctx context.Context,
	userID string,
) (*string, error) {
	f, err := os.CreateTemp("", "export-*.enc")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	defer f.Close()

	encWriter, err := q.encryptor.NewWriter(f)
	if err != nil {
		return nil, err
	}
	defer encWriter.Close()

	w := csv.NewWriter(encWriter)
	enc := csvutil.NewEncoder(w)

	for offset := int32(0); ; offset += batchSize {
		recipes, err := q.storage.ListByUserID(ctx, userID, batchSize, offset)
		if err != nil {
			return nil, err
		}

		if err := enc.Encode(recipes); err != nil {
			return nil, err
		}

		if int32(len(recipes)) < batchSize {
			break
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return nil, err
	}

	// Flush encrypted data before uploading.
	encWriter.Close()

	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s/%s.enc", userID, uuid.New().String())
	if err := q.uploader.Upload(ctx, exportBucket, key, f); err != nil {
		return nil, err
	}

	objectPath := fmt.Sprintf("%s/%s", exportBucket, key)
	return &objectPath, nil
}
