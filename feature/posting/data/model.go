package data

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
	UserID  int
	User    User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Nama     string
	Email    string `gorm:"unique"`
	Password string
	Posting  []Posting `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Posting) ToDomain() domain.Posting {
	return domain.Posting{
		ID:         int(p.ID),
		Content:    p.Content,
		Image:      p.Image,
		Created_at: p.CreatedAt,
		User: domain.UserPosting{
			ID:   int(p.UserID),
			Nama: p.User.Nama,
		},
	}
}

func ParseToArrPosting(arr []Posting) []domain.Posting {
	var res []domain.Posting
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func FromDomain(data domain.Posting) Posting {
	var res Posting
	res.Content = data.Content
	res.Image = data.Image
	res.UserID = data.User.ID
	return res
}
