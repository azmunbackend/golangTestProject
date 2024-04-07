package postgresql

import (
	"context"
	"fmt"
	"log"
	"test-crm/internal/config"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	
}

func NewClient(ctx context.Context, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	
		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			fmt.Println("failed to parse pg config: %w", err)
			// return
		}
		
		pool, err = pgxpool.ConnectConfig(ctx, cfg)
		
	if err != nil {
		fmt.Println(err)
		log.Fatal("db connectde error bar")
	}

	return pool, nil
}
