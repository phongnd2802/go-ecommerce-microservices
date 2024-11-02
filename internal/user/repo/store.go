package repo

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

// Store provides all functions to execute db queries and transactions
type sqlStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &sqlStore{
		connPool: connPool,
		Queries: New(connPool),
	}
}
