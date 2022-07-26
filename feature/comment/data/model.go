package data

import (
	"cleanarch/domain"
	"time"
)

type Comment struct {
	ID         int       `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Comment    string    `json:"comment" form:"comment"`
	Created_at time.Time `gorm:"autoCreateTime"`
	ID_Users   int
	ID_Posting int
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:         c.ID,
		Comment:    c.Comment,
		Created_at: c.Created_at,
		ID_Users:   c.ID_Users,
		ID_Posting: c.ID_Posting,
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
	res.Comment = data.Comment
	res.Created_at = data.Created_at
	res.ID_Users = data.ID_Users
	res.ID_Posting = data.ID_Posting

	return res
}
