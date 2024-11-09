package server

import (
	"context"
	"net/http"

	"github.com/geekodour/shd_24_ctf/backend/internal/db"
	"github.com/geekodour/shd_24_ctf/backend/internal/telemetry"
	"github.com/geekodour/shd_24_ctf/backend/cmd/ctf/internal/flag"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerOptions struct {
	WebPort string `default:"8000" help:"web server port"`
	Logging telemetry.LoggingConfig `embed:"" prefix:"logging."`
	// Logging struct {
	// 	Level string `enum:"debug,info,warn,error" default:"info"`
	// 	Type  string `enum:"json,console" default:"console"`
	// } `embed:"" prefix:"logging."`
	// } `embed:"" prefix:"logging."`
}
type Server struct {
	Mux    http.Handler
	Dbpool *pgxpool.Pool // pgx
	q      *db.Queries   // sqlc
	logger telemetry.Logger
	flag *flag.FlagHandler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func NewServer(config *ServerOptions) (*Server, error) {
	logger := telemetry.NewLogger(config.Logging.Level, config.Logging.Type)
	dbpool, err := db.DbPool(context.TODO(), 10, logger, config.Logging.Level)
	if err != nil {
		return nil, err
	}
	queries := db.New(dbpool)
	flagInstance := flag.NewFlagHandler(queries, dbpool)

	server := &Server{
		Dbpool: dbpool,
		q:      queries,
		logger: logger,
		flag: &flagInstance,
	}
	server.Mux = addRoutes(server, logger.Logger)
	return server, nil
}
