package encrypt

import "io"

// Encryptor は暗号化された io.WriteCloser を生成するインターフェース。
type Encryptor interface {
	NewWriter(dest io.Writer) (io.WriteCloser, error)
}
