package yakvs

import (
	"errors"
)

var (
	// ErrOpenDatastoreFile : can not open yakvs datastore file that exists.
	ErrOpenDatastoreFile = errors.New("can not open yakvs datastore file")
)
