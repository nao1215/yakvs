// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

import (
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	t.Run("Success to open yakvs", func(t *testing.T) {
		t.Parallel()

		path := filepath.Join("testdata", "testdata.yakvs")
		yakvs, err := Open(path)
		if err != nil {
			t.Errorf("expect=nil, got=error:%v", err)
		}
		yakvs.Close()
	})
}
