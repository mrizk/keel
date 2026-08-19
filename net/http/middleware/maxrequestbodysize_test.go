package middleware_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/foomo/keel/net/http/middleware"
)

func TestMaxRequestBodySize(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		hideLength bool
		wantCode   int
		wantErr    bool // handler's read of the body fails
	}{
		{name: "under the limit", body: "1234", wantCode: http.StatusOK},
		{name: "at the limit", body: "1234567890", wantCode: http.StatusOK},
		{name: "declared over the limit", body: "12345678901", wantCode: http.StatusRequestEntityTooLarge},
		{name: "undeclared over the limit", body: "12345678901", hideLength: true, wantCode: http.StatusOK, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var readErr error

			r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", strings.NewReader(tt.body))
			if tt.hideLength {
				r.ContentLength = -1
			}

			rec := httptest.NewRecorder()

			middleware.MaxRequestBodySize(10)(zap.NewNop(), "test",
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, readErr = io.ReadAll(r.Body)
				}),
			).ServeHTTP(rec, r)

			assert.Equal(t, tt.wantCode, rec.Code)

			var maxErr *http.MaxBytesError
			assert.Equal(t, tt.wantErr, errors.As(readErr, &maxErr))
		})
	}
}
