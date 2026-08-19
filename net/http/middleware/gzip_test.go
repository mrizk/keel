package middleware_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	stdhttp "github.com/foomo/gostandards/http"
	"github.com/foomo/keel/keeltest"
	"github.com/foomo/keel/net/http/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const (
	gzipPayload1023 = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. \n\nDuis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisisd"
	gzipPayload1024 = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. \n\nDuis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisisda"
)

func TestGzip(t *testing.T) {
	t.Parallel()

	svr := keeltest.NewServer(t)

	// get logger
	l := svr.Logger()

	var payload string

	svr.AddService(
		keeltest.NewServiceHTTP(l, "demo",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// validate request headers
				assert.Empty(t, r.Header.Get(stdhttp.HeaderContentEncoding.String()))
				assert.Equal(t, stdhttp.EncodingGzip.String(), r.Header.Get(stdhttp.HeaderAcceptEncoding.String()))

				// validate request body
				body, err := io.ReadAll(r.Body)
				assert.NoError(t, err)
				assert.Equal(t, payload, string(body))

				// response
				_, _ = w.Write(body)
			}),
			middleware.GZip(),
		),
	)

	svr.Start()

	url := svr.GetService("demo").URL() + "/"

	// send payload < 1024
	t.Run("<1024", func(t *testing.T) {
		payload = gzipPayload1023
		body, err := gzipString(payload)
		require.NoError(t, err)

		// send gzip request
		req, err := http.NewRequestWithContext(svr.Context(), http.MethodPost, url, bytes.NewReader(body))
		require.NoError(t, err)
		req.Header.Set(stdhttp.HeaderAcceptEncoding.String(), stdhttp.EncodingGzip.String())
		req.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()

		// validate response header
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		// validate response body
		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, payload, string(respBody))
	})

	// send payload > 1024
	t.Run(">=1024", func(t *testing.T) {
		payload = gzipPayload1024
		body, err := gzipString(payload)
		require.NoError(t, err)

		// send gzip request
		req, err := http.NewRequestWithContext(svr.Context(), http.MethodPost, url, bytes.NewReader(body))
		require.NoError(t, err)
		req.Header.Set(stdhttp.HeaderAcceptEncoding.String(), stdhttp.EncodingGzip.String())
		req.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()

		// validate response header
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, stdhttp.EncodingGzip.String(), resp.Header.Get(stdhttp.HeaderContentEncoding.String()))
		// validate response body
		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		respBodyDecompressed, err := gunzipString(respBody)
		require.NoError(t, err)
		assert.Equal(t, payload, string(respBodyDecompressed))
	})
}

func TestGZipBadRequest(t *testing.T) {
	t.Parallel()

	svr := keeltest.NewServer(t)

	// get logger
	l := svr.Logger()

	svr.AddService(
		keeltest.NewServiceHTTP(l, "demo",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			middleware.GZip(),
		),
	)

	svr.Start()

	url := svr.GetService("demo").URL() + "/"

	req, err := http.NewRequestWithContext(svr.Context(), http.MethodPost, url, strings.NewReader("hello"))
	require.NoError(t, err)

	req.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	defer resp.Body.Close()
}

func TestGZipEmptyBody(t *testing.T) {
	t.Parallel()

	handler := middleware.GZip()(zap.NewNop(), "test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// encoding/length headers must be stripped and length marked unknown
		assert.Empty(t, r.Header.Get(stdhttp.HeaderContentEncoding.String()))
		assert.Empty(t, r.Header.Get(stdhttp.HeaderContentLength.String()))
		assert.Equal(t, int64(-1), r.ContentLength)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Empty(t, body)
	}))

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", http.NoBody)
	r.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGZipContentLengthReset(t *testing.T) {
	t.Parallel()

	payload := strings.Repeat("A", 4096)
	compressed, err := gzipString(payload)
	require.NoError(t, err)

	handler := middleware.GZip()(zap.NewNop(), "test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// after decompression the compressed length no longer applies
		assert.Empty(t, r.Header.Get(stdhttp.HeaderContentLength.String()))
		assert.Equal(t, int64(-1), r.ContentLength)

		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		assert.Equal(t, payload, string(body))
	}))

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", bytes.NewReader(compressed))
	r.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGZipCaseInsensitiveEncoding(t *testing.T) {
	t.Parallel()

	payload := strings.Repeat("A", 2048)
	compressed, err := gzipString(payload)
	require.NoError(t, err)

	handler := middleware.GZip()(zap.NewNop(), "test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, readErr := io.ReadAll(r.Body)
		assert.NoError(t, readErr)
		assert.Equal(t, payload, string(body))
	}))

	r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", bytes.NewReader(compressed))
	// upper-case content-coding token must still be decompressed
	r.Header.Set(stdhttp.HeaderContentEncoding.String(), "GZIP")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGZipMaxDecompressedSize(t *testing.T) {
	t.Parallel()

	payload := strings.Repeat("A", 100*1024) // compresses tiny, expands to 100KB
	compressed, err := gzipString(payload)
	require.NoError(t, err)

	newHandler := func(limit int64) http.Handler {
		return middleware.GZip(middleware.GZipWithMaxDecompressedSize(limit))(zap.NewNop(), "test",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, readErr := io.ReadAll(r.Body)
				assert.NoError(t, readErr)
				assert.Equal(t, payload, string(body))
			}),
		)
	}

	t.Run("exceeds limit", func(t *testing.T) {
		t.Parallel()

		r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", bytes.NewReader(compressed))
		r.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

		w := httptest.NewRecorder()
		middleware.GZip(middleware.GZipWithMaxDecompressedSize(1024))(zap.NewNop(), "test",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				t.Error("handler must not be reached when limit is exceeded")
			}),
		).ServeHTTP(w, r)

		assert.Equal(t, http.StatusRequestEntityTooLarge, w.Code)
	})

	t.Run("within limit", func(t *testing.T) {
		t.Parallel()

		r := httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", bytes.NewReader(compressed))
		r.Header.Set(stdhttp.HeaderContentEncoding.String(), stdhttp.EncodingGzip.String())

		w := httptest.NewRecorder()
		newHandler(200*1024).ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func gzipString(body string) ([]byte, error) {
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(body))
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func gunzipString(body []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, gr)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
