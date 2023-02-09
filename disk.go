// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

import (
	"fmt"
	"os"

	"github.com/nao1215/gorky/file"
)

// disk is  persisting data in a file.
type disk struct {
	// file store key-value.
	file *os.File
}

// newDisk initialize disk structure.
func newDisk(path string) (*disk, error) {
	if file.IsFile(path) {
		return readYakvsDatastoreFile(path)
	}
	return createYakvsDatastoreFile(path)
}

// readYakvsDatastoreFile reads a previously created key-value store file.
func readYakvsDatastoreFile(path string) (*disk, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrOpenDatastoreFile, err)
	}

	//TODO: read meta-data
	return &disk{
		file: file,
	}, nil
}

// createYakvsDatastoreFile create new key-value store file.
func createYakvsDatastoreFile(path string) (*disk, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrOpenDatastoreFile, err)
	}

	//TODO: set meta-data
	return &disk{
		file: file,
	}, nil
}

// close close datastore file.
func (d *disk) close() error {
	return d.file.Close()
}
