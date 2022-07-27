package data

import (
	"cleanarch/domain"
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	ID        uint   `gorm:"primarykey"`
	Content   string `json:"content" form:"content"`
	Image     string `json:"image" form:"image"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Users     User           `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Nama    string
	Posting []Posting
}

func (p *Posting) ToDomain() domain.Posting {
	return domain.Posting{
		ID:         int(p.ID),
		Content:    p.Content,
		Image:      p.Image,
		Created_at: p.CreatedAt,
		User: domain.User{
			ID:   int(p.UserID),
			Nama: p.Users.Nama,
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
	res.UserID = data.ID_Users
	return res
}
