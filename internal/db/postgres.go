package db

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
    Pool *pgxpool.Pool
}

func NewPostgres(connString string) (*Database, error) {
    pool, err := pgxpool.New(context.Background(), connString)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to postgres: %w", err)
    }

    return &Database{Pool: pool}, nil
}
