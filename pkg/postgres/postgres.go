package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
)


type postgres struct {
	connPool *pgxpool.Pool
}

// Close implements DBEngine.
func (p *postgres) Close() {
	if p.GetDB() != nil {
		p.connPool.Close()
	}
}

// GetDB implements DBEngine.
func (p *postgres) GetDB() *pgxpool.Pool {
	return p.connPool
}

func NewPostgresDB(cfg settings.PostgresSetting) (DBEngine, error) {
	connPool, err := pgxpool.New(context.Background(), cfg.Addr())
	if err != nil {
		return nil, err
	}
	pg := &postgres{
		connPool: connPool,
	}
	return pg, nil
}
