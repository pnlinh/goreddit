package stores

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pnlinh/goreddit"
)

func NewCommentStore(DB *sqlx.DB) *CommentStore {
	return &CommentStore{DB: DB}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) ListByPost(postID uuid.UUID) ([]goreddit.Comment, error) {
	var comments []goreddit.Comment
	err := s.Get("SELECT * FROM comments WHERE post_id=$1", postID.String())
	if err != nil {
		return []goreddit.Comment{}, fmt.Errorf("failed to list comments: %w", err)
	}
	return comments, nil
}

func (s *CommentStore) Comment(id uuid.UUID) (goreddit.Comment, error) {
	var comment goreddit.Comment
	err := s.Get(&comment, "SELECT * FROM comments WHERE id=$1", id.String())
	if err != nil {
		return goreddit.Comment{}, fmt.Errorf("comment not found: %w", err)
	}
	return comment, nil
}

func (s *CommentStore) CreateComment(c *goreddit.Comment) error {
	err := s.Get(c, "INSERT INTO comments VALUES ($1, $2, $3, $4) RETURNING *",
		c.ID,
		c.PostID,
		c.Content,
		c.Votes,
	)
	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(c *goreddit.Comment) error {
	err := s.Get(c, "UPDATE comments SET post_id=$1, content=$2, votes=$3 WHERE id=$4 RETURNING *",
		c.PostID,
		c.Content,
		c.Votes,
		c.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update comment: %w", err)
	}
	return nil
}

func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	_, err := s.Exec("DELETE FROM comments WHERE id=$1", id.String())
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}
