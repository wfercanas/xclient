package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/selsa-inube/xclient-service/config"
)

func LogRequests(app *config.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		proto := r.Proto
		method := r.Method
		uri := r.URL.RequestURI()

		app.Logger.Info("received request", slog.String("ip", ip), slog.String("proto", proto), slog.String("method", method), slog.String("uri", uri))

		next.ServeHTTP(w, r)
	})
}
