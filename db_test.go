// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

import "testing"

func Test_dummy(t *testing.T) {
	t.Parallel()

	t.Run("dummy_test", func(t *testing.T) {
		want := "dummy"
		got := dummy()
		if want != got {
			t.Errorf("mismatch want=%s, got=%s", want, got)
		}
	})
}
