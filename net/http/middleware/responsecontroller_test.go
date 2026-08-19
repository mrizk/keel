package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/foomo/keel/net/http/middleware"
)

// deadlineRecorder captures what http.NewResponseController sets.
type deadlineRecorder struct {
	http.ResponseWriter
	write    time.Time
	writeSet bool
}

func (w *deadlineRecorder) SetWriteDeadline(t time.Time) error {
	w.write, w.writeSet = t, true
	return nil
}

func TestResponseController(t *testing.T) {
	tests := []struct {
		name     string
		h        func(rc *http.ResponseController) error
		writeSet bool
	}{
		{name: "nil handler is a no-op"},
		{
			name:     "clears the write deadline",
			h:        func(rc *http.ResponseController) error { return rc.SetWriteDeadline(time.Time{}) },
			writeSet: true,
		},
		{
			name: "error does not fail the request",
			h:    func(rc *http.ResponseController) error { return errors.New("boom") },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &deadlineRecorder{ResponseWriter: httptest.NewRecorder()}

			var called bool

			middleware.ResponseController(tt.h)(zap.NewNop(), "test",
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { called = true }),
			).ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil))

			assert.True(t, called, "handler must always run")
			assert.Equal(t, tt.writeSet, rec.writeSet)

			if tt.writeSet {
				assert.True(t, rec.write.IsZero())
			}
		})
	}
}
