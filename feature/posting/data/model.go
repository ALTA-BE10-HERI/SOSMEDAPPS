package data

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	ID_Users int
	Content  string `json:"content" form:"content"`
	Image    string `json:"image" form:"image"`
}

func (p *Posting) ToDomain() domain.Posting {
	return domain.Posting{
		ID:         int(p.ID),
		ID_Users:   p.ID_Users,
		Content:    p.Content,
		Image:      p.Image,
		Created_at: p.CreatedAt,
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
	res.ID_Users = data.ID_Users
	res.Content = data.Content
	res.Image = data.Image
	return res
}
