package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize = 10
)

var Builder squirrel.StatementBuilderType = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type tracer interface {
	pgx.QueryTracer
}

// Postgres -.
type Postgres struct {
	maxPoolSize int
	tracer      tracer

	Pool *pgxpool.Pool
}

// New -.
func New(ctx context.Context, url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize: _defaultMaxPoolSize,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed parse config: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)
	poolConfig.ConnConfig.Tracer = pg.tracer

	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	err = pg.Pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed ping db: %w", err)
	}

	return pg, nil
}

// Close -.
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
