package query

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/google/uuid"
	"github.com/jszwec/csvutil"
	"github.com/kodaikumatani/grpc-cqrs-go/internal/encrypt"
)

const batchSize int32 = 100

type Query struct {
	storage   Storage
	encryptor encrypt.Encryptor
}

func NewQuery(
	storage Storage,
	encryptor encrypt.Encryptor,
) *Query {
	return &Query{
		storage:   storage,
		encryptor: encryptor,
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

	path := f.Name()

	return &path, nil
}
