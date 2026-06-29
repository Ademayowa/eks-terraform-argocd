package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
}

var Pool Querier

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		panic("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	if err = pool.Ping(context.Background()); err != nil {
		panic("could not ping database: " + err.Error())
	}

	Pool = pool
	fmt.Println("Database connected successfully!")

	createTable()
	fmt.Println("Database tables created successfully!")
}

func Close() {
	if p, ok := Pool.(*pgxpool.Pool); ok {
		p.Close()
	}
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS events (
			id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
			title TEXT NOT NULL,
			location TEXT NOT NULL,
			date DATE NOT NULL,
			description TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`

	if _, err := Pool.Exec(context.Background(), query); err != nil {
		panic("could not create events table: " + err.Error())
	}
}
