package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/geekodour/shd_24_ctf/backend/cmd/ctf/internal/flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)


// we're checking the connection with the db as-well in our health-check, as the
// sole purpose of our application is to get data from the db, it's fine to do so.
func (s *Server) healthz(w http.ResponseWriter, req *http.Request) {
	if err := s.Dbpool.Ping(context.Background()); err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func FlagRouter(h *flag.FlagHandler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))

    // default Content-Type header for all HTTP handlers
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            next.ServeHTTP(w, r)
        })
    })

	r.Get("/get_flag/{f_id}", h.Get)
	r.Post("/submit_flag", h.Submit)
	return r
}


// Web endpoints
func addRoutes(s *Server, logger *slog.Logger) http.Handler {
	logConfig := slogchi.Config{
		WithUserAgent:      false,
		WithRequestBody:    false,
		WithResponseBody:   false,
		WithRequestHeader:  false,
		WithResponseHeader: false,
		WithRequestID:      false,
		WithSpanID:         false,
		WithTraceID:        false,
	}
	mux := chi.NewRouter()
	mux.Use(slogchi.NewWithConfig(logger, logConfig))
	mux.Use(middleware.Recoverer)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("gm"))
	})
	mux.Get("/healthz", s.healthz)

	mux.Mount("/", FlagRouter(s.flag))

	return mux
}
