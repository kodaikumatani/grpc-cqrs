package encrypt

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAESEncryptor,
)
