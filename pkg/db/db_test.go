package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxtest"

	"gophkeeper/pkg/config"
	"gophkeeper/pkg/logger"
)

func TestNew(t *testing.T) {
	t.Run("can't connect to db", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN", "invalid_dsn")
		require.NoError(t, err)

		conf, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, conf)

		log, err := logger.New(logger.Params{
			Lifecycle: fxtest.NewLifecycle(t),
			Config:    conf,
		})
		require.NoError(t, err)

		db, err := New(Params{
			Config: conf,
			Logger: log,
		})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if db != nil {
			t.Fatalf("expected db to be nil, got %v", db)
		}
	})

	t.Run("can connect to db", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		init := testInit(t)
		require.NotNil(t, init)
	})

	t.Run("ping fails", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN", "postgres://invalid_user:invalid_pass@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		conf, err := config.NewConfig()
		require.NoError(t, err)
		require.NotNil(t, conf)

		log, err := logger.New(logger.Params{
			Lifecycle: fxtest.NewLifecycle(t),
			Config:    conf,
		})
		require.NoError(t, err)

		db, err := New(Params{
			Config: conf,
			Logger: log,
		})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if db != nil {
			t.Fatalf("expected db to be nil, got %v", db)
		}
	})
}

func testInit(t *testing.T) Conn {
	conf, err := config.NewConfig()
	require.NoError(t, err)
	require.NotNil(t, conf)

	lifecycle := fxtest.NewLifecycle(t)

	log, err := logger.New(logger.Params{
		Lifecycle: lifecycle,
		Config:    conf,
	})
	require.NoError(t, err)

	db, err := New(Params{
		Lifecycle: lifecycle,
		Config:    conf,
		Logger:    log,
	})

	require.NoError(t, err)
	require.NotNil(t, db)

	return db
}

func Test_dbConn_Exec(t *testing.T) {
	t.Run("Exec", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		tag, err := db.Exec(t.Context(), "SELECT 1")
		require.NoError(t, err)
		require.NotNil(t, tag)
	})
	t.Run("Exec with error", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		// Intentionally using an invalid SQL statement to trigger an error
		_, err = db.Exec(t.Context(), "INVALID SQL STATEMENT")
		require.Error(t, err, "Expected an error for invalid SQL statement")
		require.Contains(t, err.Error(), "syntax error", "Expected syntax error in the message")
	})
}

func Test_dbConn_Query(t *testing.T) {
	t.Run("Query", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		rows, err := db.Query(t.Context(), "SELECT 1")
		require.NoError(t, err)
		require.NotNil(t, rows)

		defer rows.Close()
		for rows.Next() {
			var result int
			err = rows.Scan(&result)
			require.NoError(t, err)
			require.Equal(t, 1, result)
		}
	})

	t.Run("Query with error", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		// Intentionally using an invalid SQL statement to trigger an error
		_, err = db.Query(t.Context(), "INVALID SQL STATEMENT")
		require.Error(t, err, "Expected an error for invalid SQL statement")
		require.Contains(t, err.Error(), "syntax error", "Expected syntax error in the message")
	})
}

func Test_dbConn_QueryRow(t *testing.T) {
	t.Run("QueryRow", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		row := db.QueryRow(t.Context(), "SELECT 1")
		var result int
		err = row.Scan(&result)
		require.NoError(t, err)
		require.Equal(t, 1, result)
	})

	t.Run("QueryRow with error", func(t *testing.T) {
		err := os.Setenv("DATABASE_DSN",
			"postgres://postgres:postgres@localhost:5432/goph_keeper_db?sslmode=disable")
		require.NoError(t, err)

		db := testInit(t)
		require.NotNil(t, db)

		// Intentionally using an invalid SQL statement to trigger an error
		row := db.QueryRow(t.Context(), "INVALID SQL STATEMENT")
		var result int
		err = row.Scan(&result)
		require.Error(t, err, "Expected an error for invalid SQL statement")
	})
}
