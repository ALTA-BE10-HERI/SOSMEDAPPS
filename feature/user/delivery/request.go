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

type UpdateFormat struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UpdateFormat) ToModel() domain.User {
	return domain.User{
		ID:       u.ID,
		Nama:     u.Nama,
		Email:    u.Email,
		Password: u.Password,
	}
}
