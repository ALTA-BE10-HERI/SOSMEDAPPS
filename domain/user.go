package domain

import "cleanarch/feature/user"

type User struct {
	ID       int
	Nama     string
	Email    string
	Password string
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	GetAll() ([]User, error)
	GetProfile(id int) (User, error)
	LoginUserCase(authData user.LoginModel) (token, name string, err error)
	// DeleteCase(userID int) bool
}

type UserData interface {
	Insert(newUser User) User
	GetAll() []User
	GetSpecific(userID int) (User, error)
	LoginUserData(authData user.LoginModel) (token, name string, err error)
	// DeleteData(userID int) bool
}
