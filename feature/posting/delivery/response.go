package delivery

import (
	"cleanarch/domain"
)

type Posting struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Image   string `json:"image"`
	User    User   `json:"user"`
}

type User struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FromModel(data domain.Posting) Posting {
	return Posting{
		ID:      data.ID,
		Content: data.Content,
		Image:   data.Image,
		User: User{
			ID:   data.User.ID,
			Nama: data.User.Nama,
		},
	}
}

// func FromModelList(data []domain.Posting) []Posting {
// 	result := []Posting{}
// 	for key := range data {
// 		result = append(result, FromModel(data[key]))
// 	}
// 	return result
// }
