package healthz

import "time"

// TimeProvider is an interface that abstracts time-related functions,
// allowing for easier testing by enabling the use of mock time providers.
type TimeProvider interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

// RealTimeProvider is a struct that implements the TimeProvider interface
// using the real time functions from the time package.
type RealTimeProvider struct{}

// Now returns the current local time.
func (r *RealTimeProvider) Now() time.Time {
	return time.Now()
}

// Since returns the time elapsed since t.
func (r *RealTimeProvider) Since(t time.Time) time.Duration {
	return time.Since(t)
}
