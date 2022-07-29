package domain

import (
	"time"
)

type Comment struct {
	ID        int
	Comment   string
	Createdat time.Time
	Deletedat time.Time
	UserID    int
	PostingID int
	Nama      string
	User      UserComment
	Posting   Posting
}

type UserComment struct {
	ID   int
	Nama string
}

type CommentUseCase interface {
	CreateData(input Comment) (row int, err error)
	GetCommentByIdPosting(idPosting, limitint, offsetint int) (data []Comment, err error)
	DeleteCommentById(idComment, idFromToken int) (row int, err error)
}

type CommentData interface {
	InsertData(input Comment) (row int, err error)
	SelectCommentByIdPosting(idPosting, limitint, offsetint int) (data []Comment, err error)
	DeleteCommentByIdComment(idComment, idFromToken int) (row int, err error)
}
