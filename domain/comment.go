package domain

import (
	"time"
)

type Comment struct {
	ID         int
	Comment    string
	Created_at time.Time
	Deleted_at time.Time
	ID_Users   int
	ID_Posting int
}

type CommentUseCase interface {
	AddComment(ID_Users int, newText Comment) (Comment, error)
	GetAllComment() ([]Comment, error)
	DeleteComment(IDComment int) (row int, err error)
}

type CommentData interface {
	Insert(newText Comment) Comment
	GetComment() []Comment
	Delete(IDComment int) (row int, err error)
}
