package domain

import "time"

type Posting struct {
	ID         int
	ID_Users   int
	Content    string
	Image      string
	Created_at time.Time
	Deleted_at time.Time
}

type PostingUserCase interface {
	AddPosting(userID int, newPosting Posting) (Posting, error)
	GetAllPosting() ([]Posting, error)
	// UpdateCase(updatePost Posting) (Posting, error)
	// DeleteCase(ID int) error
}

type PostingData interface {
	Insert(newPosting Posting) Posting
	GetPosting() []Posting
	// UpdatePost(updatePost Posting) Posting
	// DeletePost(ID int) error
}
