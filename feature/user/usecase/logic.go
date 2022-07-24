package usecase

import (
	"cleanarch/domain"
	"cleanarch/feature/user"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(uc domain.UserData) domain.UserUseCase {
	return &userUseCase{
		userData: uc,
	}
}

func (uc *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted := uc.userData.Insert(newUser)
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}

func (uc *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := uc.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (uc *userUseCase) LoginUserCase(authData user.LoginModel) (token, name string, err error) {
	token, name, err = uc.userData.LoginUserData(authData)
	return token, name, err
}

func (uc *userUseCase) DeleteCase(userID int) (row int, err error) {
	row, err = uc.userData.DeleteData(userID)
	return row, err
}

func (uc *userUseCase) UpdateCase(id int, newUser domain.User) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted := uc.userData.UpdateData(newUser, newUser)
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}
