package delivery

import (
	"cleanarch/domain"
	"time"
)

type Comment struct {
	ID        int `json:"id"`
	Nama      string
	Comment   string `json:"comment"`
	Createdat time.Time
}

func FromModel(data domain.Comment) Comment {
	return Comment{
		ID:        data.ID,
		Nama:      data.User.Nama,
		Comment:   data.Comment,
		Createdat: data.Createdat,
	}
}

func FromModelList(data []domain.Comment) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
