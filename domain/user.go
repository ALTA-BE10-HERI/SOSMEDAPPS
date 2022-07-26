package domain

import (
	"cleanarch/feature/user"
	"time"
)

type User struct {
	ID        int
	Nama      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	GetProfile(id int) (User, error)
	LoginUserCase(authData user.LoginModel) (token, name string, err error)
	DeleteCase(userID int) (row int, err error)

	UpdateCase(input User, idUser int) (row int, err error)
}

type UserData interface {
	Insert(newUser User) User
	GetSpecific(userID int) (User, error)
	LoginUserData(authData user.LoginModel) (token, name string, err error)
	DeleteData(userID int) (row int, err error)

	UpdateData(data map[string]interface{}, idUser int) (row int, err error)
}
