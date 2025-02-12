package middleware

import (
	"net/http"
	"time"

	"kai_security/internal/logger"

	"github.com/rs/zerolog/hlog"
)

func requestLogger(next http.Handler) http.Handler {
	l := logger.Get()

	h := hlog.NewHandler(l)

	requestIdhandler := hlog.RequestIDHandler("X-Correlation-Id", "X-Correlation-Id")

	accessHandler := hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("status_code", status).
				Int("response_size_bytes", size).
				Dur("elapsed_ms", duration).
				Msg("incoming request")
		},
	)

	userAgentHandler := hlog.UserAgentHandler("http_user_agent")

	return h(requestIdhandler(accessHandler(userAgentHandler(next))))
}
