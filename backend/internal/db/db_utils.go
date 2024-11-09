package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"os"
)

// NOTE: NOTIFY/LISTEN
//   - NOTIFY/LISTEN doesn't work with Transaction Pooling, only with Session Pooling
//   - LISTEN is tied to a client connection, so if you are using a connection pool
//     on a transaction level (i.e. when you finish a transaction, your real connection
//     is return to the pool) you can run into trouble.
//   - https://github.com/pgbouncer/pgbouncer/issues/655
//   - So whenever we want to use LISTEN, we should always use a client side pool of size 1.
//
//	Since we're not using PG bouncer at the moment, we don't have to worry.
//	But if we decide to use pgbouncer, we'll need to change this method to
//	connect directly to postgres for LISTEN as pgbouncer has issues with LISTEN
//
//  Instead of pool, we could get a direct connection aswell. For now going with
//	poolsize=1 for the same effect.
// func LISTENPool(ctx context.Context, logger tracelog.Logger) (*pgxpool.Pool, error) {
// 	return DbPool(ctx, 1, logger, "")
// }

// NOTE: tracing support for tracelog(pgx)
//
//	https://github.com/jackc/pgx/discussions/1677#discussioncomment-8815982
//
// NOTE: TraceLog doesn't take logLevel as we have in our logging library it
//
//		 takes in a number, we will need to add a mapping fucntion, for now
//	 	 we're just keeping the argument for nothing but its OK
func DbPool(ctx context.Context, maxConnection int32, logger tracelog.Logger, logLevel string) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}
	// See https://pkg.go.dev/github.com/jackc/pgx/v5#hdr-PgBouncer
	// NOTE: We disabling QueryExecModeSimpleProtocol because of https://github.com/sqlc-dev/sqlc/issues/2085
	// connConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	connConfig.ConnConfig.Tracer = &tracelog.TraceLog{Logger: logger, LogLevel: 5}
	connConfig.MaxConns = maxConnection

	dbpool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
