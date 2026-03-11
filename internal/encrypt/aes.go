package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
)

// AESEncryptor は AES-CTR モードで暗号化を行う Encryptor の実装。
type AESEncryptor struct{}

func NewAESEncryptor() Encryptor {
	return &AESEncryptor{}
}

func (e *AESEncryptor) NewWriter(dest io.Writer) (io.WriteCloser, error) {
	keyHex, err := hex.DecodeString(os.Getenv("AES_KEY"))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(keyHex)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	if _, err := dest.Write(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	return &cipher.StreamWriter{S: stream, W: dest}, nil
}
