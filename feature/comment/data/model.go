package data

import (
	"cleanarch/domain"
	"time"
)

type Comment struct {
	ID         int `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	ID_Posting int
	ID_Users   int
	Comment    string    `json:"comment" form:"comment"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Deleted_at time.Time `gorm:"autoCreateTime"`
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:         int(c.ID),
		ID_Posting: c.ID_Posting,
		ID_Users:   c.ID_Users,
		Comment:    c.Comment,
		Created_at: c.Created_at,
		Deleted_at: c.Deleted_at,
	}
}

func ParseToArrComment(arr []Comment) []domain.Comment {
	var res []domain.Comment
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func FromDomain(data domain.Comment) Comment {
	var res Comment
	res.ID = data.ID
	res.ID_Posting = data.ID_Posting
	res.ID_Users = data.ID_Users
	res.Comment = data.Comment
	return res
}
