package yakvs

import (
	"errors"
	"fmt"
)

// Wrap wrap error message.
func Wrap(ep *error, format string, args ...any) {
	if *ep != nil {
		*ep = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *ep)
	}
}

var (
	// ErrDatastoreFile : can not open yakvs datastore file that exists.
	ErrDatastoreFile = errors.New("can not open yakvs datastore file")
)
