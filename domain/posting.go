package domain

import "time"

type Posting struct {
	ID         int
	Content    string
	Image      string
	Created_at time.Time
	Deleted_at time.Time
	ID_Users   int
	User       User
}

type UserPosting struct {
	ID   int
	Name string
}

type PostingUseCase interface {
	// AddPosting(userID int, newPosting Posting) (Posting, error)
	AddPosting(data Posting) (row int, err error)
	GetAllPosting() ([]Posting, error)
	DeleteCase(postingID int) (row int, err error)
	// UpdateCase(updatePost Posting) (Posting, error)
}

type PostingData interface {
	// Insert(newPosting Posting) Posting
	InsertData(data Posting) (row int, err error)
	GetPosting() []Posting
	DeleteData(postingID int) (row int, err error)
	// UpdatePost(updatePost Posting) Posting
}
