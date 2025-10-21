package stores

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pnlinh/goreddit"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) ListByThread(threadID uuid.UUID) ([]goreddit.Post, error) {
	var posts []goreddit.Post
	err := s.Select(&posts, "SELECT * FROM posts WHERE thread_id=$1", threadID)
	if err != nil {
		return []goreddit.Post{}, fmt.Errorf("failed to list posts: %w", err)
	}
	return posts, nil
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var p goreddit.Post
	err := s.Get(&p, "SELECT * FROM posts WHERE id=$1", id)
	if err != nil {
		return goreddit.Post{}, fmt.Errorf("post not found: %w", err)
	}
	return p, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	err := s.Get(p, "INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *",
		p.ID,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
	)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	err := s.Get(p, "UPDATE posts SET thread_id=$1 title=$2, content=$3, votes=$4 WHERE id=$5 RETURNING *",
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
		p.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	_, err := s.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}
