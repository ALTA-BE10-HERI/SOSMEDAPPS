package domain

import (
	"time"
)

type Comment struct {
	ID         int
	ID_Posting int
	ID_Users   int
	Comment    string
	Created_at time.Time
	Deleted_at time.Time
}

type CommentUseCase interface {
	AddComment(newText Comment) (Comment, error)
	GetAllComment() ([]Comment, error)
	//DeleteComment(int) error
}

type CommentData interface {
	Insert(newText Comment) Comment
	GetComment() []Comment
	//DeleteComment(IDComment int) bool
}
