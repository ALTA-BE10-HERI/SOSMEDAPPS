package data

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID_Users   int
	ID_Posting int
	Comment    string `json:"comment" form:"comment"`
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:         int(c.ID),
		ID_Users:   c.ID_Users,
		ID_Posting: c.ID_Posting,
		Comment:    c.Comment,
		Created_at: c.CreatedAt,
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
	res.ID_Users = data.ID_Users
	res.ID_Posting = data.ID_Posting
	res.Comment = data.Comment
	return res
}
