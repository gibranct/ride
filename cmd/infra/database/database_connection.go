package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnection interface {
	QueryWithContext(ctx context.Context, dest any, query string, params ...any) error
	ExecContext(ctx context.Context, query string, params ...any) error
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

func (pg *PostgresAdapter) Close() error {
	err := pg.db.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to close database connection: %v\n", err)
		return err
	}
	return nil
}

func NewPostgresAdapter() *PostgresAdapter {
	db, err := sqlx.Connect("postgres", "postgres://postgres:123456@localhost:5432/app?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PostgresAdapter{
		db: db,
	}
}
