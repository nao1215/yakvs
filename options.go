// Package yakvs is Yet Another Key Value Store in golang.
package yakvs

// Option applies options when initializing YAKVS.
type Option interface {
	Apply() error
}
