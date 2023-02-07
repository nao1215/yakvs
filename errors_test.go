package yakvs

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	t.Parallel()

	t.Run("Wrap error message with \\%d format", func(t *testing.T) {
		e := errors.New("original message")
		format := "test(%d)"
		Wrap(&e, format, 1)

		want := "test(1): original message"
		got := e.Error()

		if want != got {
			t.Errorf("mismatch want='%s', got='%s'", want, got)
		}
	})

	t.Run("Wrap error message with \\%s format", func(t *testing.T) {
		e := errors.New("original message")
		format := "test(%s)"
		Wrap(&e, format, "wrap")

		want := "test(wrap): original message"
		got := e.Error()

		if want != got {
			t.Errorf("mismatch want='%s', got='%s'", want, got)
		}
	})
}
