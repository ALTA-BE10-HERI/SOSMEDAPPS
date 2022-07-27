package delivery

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type InsertFormat struct {
	gorm.Model
	Content  string `json:"content" form:"content"`
	Image    string `json:"image" form:"content"`
	ID_Users int
}

// func ToModel(i InsertFormat) domain.Posting {
// 	return domain.Posting{
// 		Content: i.Content,
// 		Image:   i.Image,
// 		User: domain.UserPosting{
// 			ID: i.ID_Users,
// 		},
// 	}
// }
func (i *InsertFormat) ToModel() domain.Posting {
	return domain.Posting{
		Content: i.Content,
		Image:   i.Image,
		User: domain.UserPosting{
			ID: i.ID_Users,
		},
	}
}
