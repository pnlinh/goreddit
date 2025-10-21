package stores

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Store{
		ThreadStore:  &ThreadStore{db},
		PostStore:    &PostStore{db},
		CommentStore: &CommentStore{db},
	}, nil
}

type Store struct {
	*ThreadStore
	*PostStore
	*CommentStore
}
