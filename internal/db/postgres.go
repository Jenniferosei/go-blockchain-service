// package db

// import (
//     "context"
//     "fmt"
//     "os"
//     "time"

//     "github.com/jackc/pgx/v5/pgxpool"
// )

// var DB *pgxpool.Pool

// func Connect() error {
//     databaseUrl := os.Getenv("DATABASE_URL")
//     if databaseUrl == "" {
//         return fmt.Errorf("DATABASE_URL environment variable not set")
//     }

//     config, err := pgxpool.ParseConfig(databaseUrl)
//     if err != nil {
//         return fmt.Errorf("unable to parse DATABASE_URL: %w", err)
//     }

//     pool, err := pgxpool.NewWithConfig(context.Background(), config)
//     if err != nil {
//         return fmt.Errorf("unable to create connection pool: %w", err)
//     }

//     // Test connection with timeout
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     err = pool.Ping(ctx)
//     if err != nil {
//         return fmt.Errorf("unable to ping database: %w", err)
//     }

//     DB = pool
//     return nil
// }

package db

import (
    "context"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(pool *pgxpool.Pool) {
    DB = pool
}

func NewPostgres(url string) (*pgxpool.Pool, error) {
    config, err := pgxpool.ParseConfig(url)
    if err != nil {
        return nil, err
    }

    config.MaxConns = 10
    config.MinConns = 2
    config.HealthCheckPeriod = 30 * time.Second
    config.MaxConnLifetime = time.Hour
    config.MaxConnIdleTime = 10 * time.Minute

    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, err
    }

    // Test connection
    if err := pool.Ping(context.Background()); err != nil {
        return nil, err
    }

    return pool, nil
}