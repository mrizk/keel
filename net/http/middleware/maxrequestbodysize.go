package middleware

import (
	"net/http"

	"go.uber.org/zap"

	keelhttp "github.com/foomo/keel/net/http"
)

// MaxRequestBodySize middleware caps the request body at n bytes.
//
// A declared Content-Length over the limit is rejected before the handler runs;
// otherwise the body is wrapped with http.MaxBytesReader, so a chunked or lying
// client fails on read with *http.MaxBytesError once it exceeds n. Handlers that
// read the body should surface that as 413:
//
//	var maxErr *http.MaxBytesError
//	if errors.As(err, &maxErr) { ... }
//
// This bounds size; it does not bound time. A slow client trickling bytes under
// the limit is the server's ReadTimeout's problem — see ResponseController to
// adjust that per request.
func MaxRequestBodySize(n int64) keelhttp.Middleware {
	return func(l *zap.Logger, name string, next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > n {
				http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
				return
			}

			if r.Body != nil {
				r.Body = http.MaxBytesReader(w, r.Body, n)
			}

			next.ServeHTTP(w, r)
		})
	}
}
