package delivery

import "cleanarch/domain"

type InsertFormat struct {
	Nama     string `json:"nama"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" `
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Nama:     i.Nama,
		Email:    i.Email,
		Password: i.Password,
	}
}
