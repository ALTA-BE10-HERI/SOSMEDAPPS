package delivery

import (
	"cleanarch/domain"
)

type Comment struct {
	ID      int `json:"id"`
	Nama    string
	Comment string `json:"comment"`
}

func FromModel(data domain.Comment) Comment {
	return Comment{
		ID:      data.ID,
		Nama:    data.Nama,
		Comment: data.Comment,
	}
}

func FromModelList(data []domain.Comment) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
