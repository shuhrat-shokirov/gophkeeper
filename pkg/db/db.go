package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	fx.Lifecycle

	Config config.Config
	Logger logger.Logger
}

type dbConn struct {
	config config.Config
	dbPool *pgxpool.Pool
	logger logger.Logger
}

type Conn interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

func New(params Params) (Conn, error) {
	var (
		dns = params.Config.GetString("database.dsn")
		err error
	)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	db, err := pgxpool.New(ctx, dns)
	if err != nil {
		params.Logger.Error(fmt.Sprintf("Err on pgxpool.Connect(%v): %v", dns, err))
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		params.Logger.Error(fmt.Sprintf("Err on db.Ping(): %v", err))
		return nil, err
	}

	params.Logger.Info("Connected to database")

	params.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				db.Close()
				params.Logger.Info("Closing database connection")
				return nil
			},
		},
	)

	return &dbConn{
		dbPool: db,
		logger: params.Logger,
		config: params.Config,
	}, nil
}

func (db *dbConn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return db.dbPool.Exec(ctx, sql, arguments...)
}

func (db *dbConn) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	return db.dbPool.Query(ctx, sql, optionsAndArgs...)
}

func (db *dbConn) QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row {
	return db.dbPool.QueryRow(ctx, sql, optionsAndArgs...)
}
