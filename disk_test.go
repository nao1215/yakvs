// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

import (
	"errors"
	"path/filepath"
	"testing"
)

func Test_newDisk(t *testing.T) {
	t.Run("open new datastore file", func(t *testing.T) {
		t.Parallel()
		filePath := filepath.Join(t.TempDir(), "new-datastore.yakvs")

		disk, err := newDisk(filePath)
		if err != nil {
			t.Errorf("expect=nil, got=error:%v", err)
		}
		disk.close()
	})

	t.Run("open previously created datastore file", func(t *testing.T) {
		t.Parallel()
		filePath := filepath.Join("testdata", "testdata.yakvs")

		disk, err := newDisk(filePath)
		if err != nil {
			t.Errorf("expect=nil, got=error:%v", err)
		}
		disk.close()
	})
}

func Test_readYakvsDatastoreFile(t *testing.T) {
	t.Run("can not read datastore file", func(t *testing.T) {
		t.Parallel()

		want := ErrOpenDatastoreFile
		_, got := readYakvsDatastoreFile("")
		if !errors.Is(got, want) {
			t.Errorf("mismatch want=%v, got=%v", want, got)
		}
	})
}

func Test_createYakvsDatastoreFile(t *testing.T) {
	t.Run("can not open datastore file", func(t *testing.T) {
		t.Parallel()

		want := ErrOpenDatastoreFile
		_, got := readYakvsDatastoreFile("")
		if !errors.Is(got, want) {
			t.Errorf("mismatch want=%v, got=%v", want, got)
		}
	})
}
