package domain

import (
	"time"
)

type Comment struct {
	ID         int
	ID_Users   int
	ID_Posting int
	Comment    string
	Created_at time.Time
	Deleted_at time.Time
	User       UserComment
	Posting    PostingComment
}

type UserComment struct {
	ID   int
	Nama string
}

type PostingComment struct {
	ID      int
	Content string
	Image   string
}

type CommentUseCase interface {
	AddComment(data Comment) (result Comment, err error)
	GetAllComment() ([]Comment, error)
	DeleteComment(IDComment int) (row int, err error)
}

type CommentData interface {
	Insert(data Comment) (result Comment, err error)
	GetComment() []Comment
	Delete(IDComment int) (row int, err error)
}
