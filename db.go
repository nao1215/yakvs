// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

// YAKVS is a structure to control yakvs (No SQL) datastore.
type YAKVS struct {
	// disk is yakvs datastore file.
	disk *disk
}

// Open open yakvs datastore file.
func Open(path string, options ...Option) (*YAKVS, error) {
	disk, err := newDisk(path)
	if err != nil {
		return nil, err
	}

	return &YAKVS{
		disk: disk,
	}, nil
}

// Close close yakvs datastore file.
func (y *YAKVS) Close() error {
	return y.disk.close()
}
