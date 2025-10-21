package goreddit

import "github.com/google/uuid"

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type ThreadStore interface {
	List() ([]*Thread, error)
	Show(id uuid.UUID) (*Thread, error)
	Create(t *Thread) error
	Update(t *Thread) error
	Delete(id uuid.UUID) error
}

type PostStore interface {
	ListByThread(threadID uuid.UUID) ([]*Post, error)
	Show(id uuid.UUID) (*Post, error)
	Create(p *Post) error
	Update(p *Post) error
	Delete(id uuid.UUID) error
}

type CommentStore interface {
	ListByPost(postID uuid.UUID) ([]*Comment, error)
	Show(id uuid.UUID) (*Comment, error)
	Create(c *Comment) error
	Update(c *Comment) error
	Delete(id uuid.UUID) error
}
