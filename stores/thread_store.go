package stores

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pnlinh/goreddit"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) List() ([]goreddit.Thread, error) {
	var threads []goreddit.Thread
	err := s.Select(&threads, "SELECT * FROM threads")
	if err != nil {
		return []goreddit.Thread{}, fmt.Errorf("failed to list threads: %w", err)
	}
	return threads, nil
}

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	err := s.Get(&t, "SELECT * FROM threads WHERE id=$1", id)
	if err != nil {
		return goreddit.Thread{}, fmt.Errorf("thread not found: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	err := s.Get(t, "INSERT INTO threads VALUES ($1, $2, $3) RETURNING *",
		t.ID,
		t.Title,
		t.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to create thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	err := s.Get(t, "UPDATE threads SET title=$1, description=$2 WHERE id=$3 RETURNING *",
		t.Title,
		t.Description,
		t.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	_, err := s.Exec("DELETE FROM threads WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to delete thread: %w", err)
	}
	return nil
}
