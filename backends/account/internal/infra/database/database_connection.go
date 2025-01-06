package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnection interface {
	QueryWithContext(ctx context.Context, dest any, query string, params ...any) error
	ExecContext(ctx context.Context, query string, params ...any) error
	SelectContext(ctx context.Context, dest any, query string, params ...any) error
	Close() error
}

type PostgresAdapter struct {
	db *sqlx.DB
}

func (pg *PostgresAdapter) QueryWithContext(ctx context.Context, dest any, query string, params ...any) error {
	return pg.db.GetContext(ctx, dest, query, params...)
}

func (pg *PostgresAdapter) ExecContext(ctx context.Context, query string, params ...any) error {
	_, err := pg.db.ExecContext(ctx, query, params...)
	return err
}

func (pg *PostgresAdapter) SelectContext(ctx context.Context, dest any, query string, params ...any) error {
	return pg.db.SelectContext(ctx, dest, query, params...)
}

func (pg *PostgresAdapter) Close() error {
	err := pg.db.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to close database connection: %v\n", err)
		return err
	}
	return nil
}

var instance *PostgresAdapter

var once sync.Once

func NewPostgresAdapter() *PostgresAdapter {
	once.Do(func() {
		connString := os.Getenv("POSTGRES_DSN")
		if connString == "" {
			connString = "postgres://postgres:123456@localhost:5433/app?sslmode=disable"
		}
		db, err := sqlx.Connect("postgres", connString)
		if err != nil {
			log.Fatalln(err)
		}
		instance = &PostgresAdapter{
			db: db,
		}
	})

	return instance
}
