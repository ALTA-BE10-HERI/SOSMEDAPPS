package delivery

import "cleanarch/domain"

type Comment struct {
	ID      int     `json:"id"`
	Comment string  `json:"comment"`
	User    User    `json:"user"`
	Posting Posting `json:"posting"`
}

type User struct {
	ID int `json:"id"`
}

type Posting struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func FromModel(data domain.Comment) Comment {
	return Comment{
		ID:      data.ID,
		Comment: data.Comment,
		User:    data.User.ID,
		Posting: Posting{
			ID:      data.Posting.ID,
			Content: data.Posting.Content,
			Image:   data.Posting.Content,
		},
	}
}

// func FromModelList(data []domain.Comment) []Comment {
// 	result := []Comment{}
// 	for key := range data {
// 		result = append(result, FromModel(data[key]))
// 	}
// 	return result
// }
