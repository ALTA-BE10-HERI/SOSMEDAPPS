package delivery

import "cleanarch/domain"

type InsertFormat struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Nama:     i.Nama,
		Email:    i.Email,
		Password: i.Password,
	}
}
