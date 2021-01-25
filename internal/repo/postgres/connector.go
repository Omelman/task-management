package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Omelman/task-management/internal/config"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Postgres struct {
	db *sql.DB
}

var (
	postgres *Postgres
	once     = &sync.Once{}
)

func Load(opts config.Postgres) error {
	once.Do(func() {
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			opts.Host, opts.Port, opts.User,
			opts.Password, opts.Database, "disable",
		)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		postgres = &Postgres{
			db: db,
		}
	})

	return postgres.db.Ping()
}

func GetDB() *Postgres {
	return postgres
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}
