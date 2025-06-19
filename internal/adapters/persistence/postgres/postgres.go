package postgres

import (
	"context"

	_ "github.com/lib/pq"
	"goraft/ent"
	"goraft/internal/core/port/outbound"
)

type PostgresClient struct {
	client *ent.Client
}

func (p *PostgresClient) Connect(ctx context.Context) error {
	c, err := ent.Open("postgres" /* DSN */, "")
	if err != nil {
		return err
	}
	p.client = c
	return p.client.Schema.Create(ctx)
}

func (p *PostgresClient) Close() error {
	return p.client.Close()
}

// compile-time check
var _ outbound.DatabaseClient = (*PostgresClient)(nil)
