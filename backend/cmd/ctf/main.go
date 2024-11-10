package main

import (
	"context"
	"fmt"
	// "fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/geekodour/shd_24_ctf/backend/cmd/ctf/internal/server"
	"github.com/geekodour/shd_24_ctf/backend/internal/telemetry"
	"github.com/oklog/run"
)

func cliCommand() *server.ServerOptions {
	opts := server.ServerOptions{}
	kong.Parse(&opts,
		kong.Name("ctf"),
		kong.Description("shd_24_ctf server"),
		kong.ShortUsageOnError(),
	)
	return &opts
}

func main() {
	// app context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srvOpts := cliCommand()

	// default logger
	// NOTE: We need to set the deault logger separately from all the other
	// 		 loggers that we setup
	logger := telemetry.NewSlogLogger(srvOpts.Logging.Level, srvOpts.Logging.Type)
	slog.SetDefault(logger)

	// define server
	srv, err := server.NewServer(srvOpts)
	if err != nil {
		slog.Info("couldn't setup server")
		return
	}
	defer srv.Dbpool.Close()

	// define channels

	// NOTE: We're running multiple services here, the idea is similar to erlang
	//		 supervisor trees. See https://www.jerf.org/iri/post/2930/
	//
	//		 But instead of using https://github.com/thejerf/suture, we're using
	//		 oklog/run as suture doesn't seem to handle net/http (doesn't fit in
	//		 the model). The idea is similar to errgroup package aswell.
	var g run.Group
	{
		g.Add(run.SignalHandler(ctx, os.Interrupt))
	}
	{
		server := &http.Server{
			Handler: srv.Mux,
			Addr:    ":" + srvOpts.WebPort,
		}
		g.Add(func() error {
			slog.Info(fmt.Sprintf("Web server listening on %s\n", srvOpts.WebPort))
			return server.ListenAndServe()
		}, func(err error) {
			slog.Info("attempting server shutdown")
			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			_ = server.Shutdown(ctx)
		})
	}

	slog.Error("exited", "reason", g.Run())
}
