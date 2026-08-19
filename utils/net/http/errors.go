package httputils

import (
	"errors"
	"net/http"

	httplog "github.com/foomo/keel/net/http/log"
	"github.com/foomo/keel/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.uber.org/zap"

	"github.com/foomo/keel/log"
)

// InternalServerError http response
func InternalServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusInternalServerError, errs...)
}

// InternalServiceUnavailable http response
func InternalServiceUnavailable(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusServiceUnavailable, errs...)
}

// InternalServiceTooEarly http response
func InternalServiceTooEarly(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusTooEarly, errs...)
}

// UnauthorizedServerError http response
func UnauthorizedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusUnauthorized, errs...)
}

// BadRequestServerError http response
func BadRequestServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusBadRequest, errs...)
}

// MethodNotAllowedServerError http response
func MethodNotAllowedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusMethodNotAllowed, errs...)
}

// RequestEntityTooLargeServerError http response
func RequestEntityTooLargeServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusRequestEntityTooLarge, errs...)
}

// NotFoundServerError http response
func NotFoundServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusNotFound, errs...)
}

// PaymentRequiredServerError http response
func PaymentRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusPaymentRequired, errs...)
}

// ForbiddenServerError http response
func ForbiddenServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusForbidden, errs...)
}

// NotAcceptableServerError http response
func NotAcceptableServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusNotAcceptable, errs...)
}

// ProxyAuthRequiredServerError http response
func ProxyAuthRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusProxyAuthRequired, errs...)
}

// RequestTimeoutServerError http response
func RequestTimeoutServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusRequestTimeout, errs...)
}

// ConflictServerError http response
func ConflictServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusConflict, errs...)
}

// GoneServerError http response
func GoneServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusGone, errs...)
}

// LengthRequiredServerError http response
func LengthRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusLengthRequired, errs...)
}

// PreconditionFailedServerError http response
func PreconditionFailedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusPreconditionFailed, errs...)
}

// RequestURITooLongServerError http response
func RequestURITooLongServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusRequestURITooLong, errs...)
}

// UnsupportedMediaTypeServerError http response
func UnsupportedMediaTypeServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusUnsupportedMediaType, errs...)
}

// RequestedRangeNotSatisfiableServerError http response
func RequestedRangeNotSatisfiableServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusRequestedRangeNotSatisfiable, errs...)
}

// ExpectationFailedServerError http response
func ExpectationFailedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusExpectationFailed, errs...)
}

// TeapotServerError http response
func TeapotServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusTeapot, errs...)
}

// MisdirectedRequestServerError http response
func MisdirectedRequestServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusMisdirectedRequest, errs...)
}

// UnprocessableEntityServerError http response
func UnprocessableEntityServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusUnprocessableEntity, errs...)
}

// LockedServerError http response
func LockedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusLocked, errs...)
}

// FailedDependencyServerError http response
func FailedDependencyServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusFailedDependency, errs...)
}

// UpgradeRequiredServerError http response
func UpgradeRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusUpgradeRequired, errs...)
}

// PreconditionRequiredServerError http response
func PreconditionRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusPreconditionRequired, errs...)
}

// TooManyRequestsServerError http response
func TooManyRequestsServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusTooManyRequests, errs...)
}

// RequestHeaderFieldsTooLargeServerError http response
func RequestHeaderFieldsTooLargeServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusRequestHeaderFieldsTooLarge, errs...)
}

// UnavailableForLegalReasonsServerError http response
func UnavailableForLegalReasonsServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusUnavailableForLegalReasons, errs...)
}

// NotImplementedServerError http response
func NotImplementedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusNotImplemented, errs...)
}

// BadGatewayServerError http response
func BadGatewayServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusBadGateway, errs...)
}

// GatewayTimeoutServerError http response
func GatewayTimeoutServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusGatewayTimeout, errs...)
}

// HTTPVersionNotSupportedServerError http response
func HTTPVersionNotSupportedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusHTTPVersionNotSupported, errs...)
}

// VariantAlsoNegotiatesServerError http response
func VariantAlsoNegotiatesServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusVariantAlsoNegotiates, errs...)
}

// InsufficientStorageServerError http response
func InsufficientStorageServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusInsufficientStorage, errs...)
}

// LoopDetectedServerError http response
func LoopDetectedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusLoopDetected, errs...)
}

// NotExtendedServerError http response
func NotExtendedServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusNotExtended, errs...)
}

// NetworkAuthenticationRequiredServerError http response
func NetworkAuthenticationRequiredServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, errs ...error) {
	ServerError(l, w, r, http.StatusNetworkAuthenticationRequired, errs...)
}

// ServerError http response
func ServerError(l *zap.Logger, w http.ResponseWriter, r *http.Request, code int, errs ...error) {
	if err := errors.Join(errs...); err != nil {
		errType := semconv.ErrorType(err)
		telemetry.Ctx(r.Context()).RecordError(err)

		if labeler, ok := otelhttp.LabelerFromContext(r.Context()); ok {
			labeler.Add(errType)
		}

		// add log entry
		if labeler, ok := httplog.LabelerFromRequest(r); ok {
			labeler.Add(log.Attribute(errType), log.FError(err))
		} else if l != nil {
			l = log.WithError(l, err)
			l = log.WithHTTPRequest(l, r)
			l.Error("http server error", log.Attribute(semconv.HTTPResponseStatusCode(code)))
		}

		http.Error(w, http.StatusText(code), code)
	}
}
