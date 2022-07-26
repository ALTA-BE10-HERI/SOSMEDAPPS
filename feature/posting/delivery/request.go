package delivery

import "cleanarch/domain"

type InsertFormat struct {
	ID_Users int    `json:"id"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}

func (i *InsertFormat) ToModel() domain.Posting {
	return domain.Posting{
		ID_Users: i.ID_Users,
		Content:  i.Content,
		Image:    i.Image,
	}
}
