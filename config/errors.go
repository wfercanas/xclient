package config

import (
	"log/slog"
	"net/http"
)

func (app *Application) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	app.Logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()), slog.Int("status", code))
	http.Error(w, err.Error(), code)
}
