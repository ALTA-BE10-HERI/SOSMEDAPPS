package domain

import "time"

type Posting struct {
	ID         int
	Content    string
	Image      string
	Created_at time.Time
	Updated_at time.Time
	User       UserPosting
}

type UserPosting struct {
	ID   int
	Nama string
}

//kirim ke logic
type PostingUseCase interface {
	AddPosting(data Posting) (result Posting, err error)
	GetAllPosting() ([]Posting, error)
	DeleteCase(idPosting, idFromToken int) (row int, err error)
	GetPostingById(id int) (data Posting, err error)
	UpdateData(data Posting, idPosting, idFromToken int) (row int, err error)
}

//kirim ke query
type PostingData interface {
	GetUser(idPosting int) (result Posting, err error) //buat ambil nama id
	InsertData(data Posting) (result Posting, err error)
	GetPosting() []Posting
	DeleteDataById(idPosting, idFromToken int) (row int, err error)
	SelectDataById(id int) (data Posting, err error)
	UpdateData(data map[string]interface{}, idPosting, idFromToken int) (res int, err error)
}
