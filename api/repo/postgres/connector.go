package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Omelman/task-management/api/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	db *sql.DB
}

func (db *DB) RawDB() *sql.DB {
	return db.db
}

func (db *DB) Clone() *DB {
	return &DB{
		db: db.db,
	}
}

func Open(opts config.Postgres) (*DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		opts.Host, opts.Port, opts.User,
		opts.Password, opts.Database, "disable",
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return &DB{
		db: db,
	}, nil
}
