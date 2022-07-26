package data

import (
	"cleanarch/domain"
	"cleanarch/feature/posting/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string `gorm:"unique" validate:"required,email"`
	Password string
	Posting  []data.Posting `gorm:"foreignKey:ID_Users"` // masih belum yakin
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		Nama:      u.Nama,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User
	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Email = data.Email
	res.Nama = data.Nama
	res.Password = data.Password
	return res
}
