package data

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostingID int
	UserID    int
	Comment   string  `json:"comment" form:"comment"`
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Posting   Posting `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Nama    string
	Comment []Comment
}

type Posting struct {
	gorm.Model
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		PostingID: int(c.PostingID),
		ID:        int(c.ID),
		Comment:   c.Comment,
		Createdat: c.CreatedAt,
		UserID:    int(c.UserID),
		User: domain.UserComment{
			ID:   int(c.User.ID),
			Nama: c.User.Nama,
		},
	}
}

func parseToArrComment(arr []Comment) []domain.Comment {
	var res []domain.Comment
	for val := range arr {
		res = append(res, arr[val].ToDomain())
	}

	return res
}

func FromDomain(data domain.Comment) Comment {
	var res Comment
	res.UserID = data.UserID
	res.PostingID = data.PostingID
	res.Comment = data.Comment
	return res
}
