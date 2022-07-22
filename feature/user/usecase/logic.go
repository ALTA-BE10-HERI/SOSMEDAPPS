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

func New(ud domain.UserData) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
	}
}

func (ud *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted := ud.userData.Insert(newUser)
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}

func (ud *userUseCase) GetAll() ([]domain.User, error) {
	data := ud.userData.GetAll()

	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	return data, nil
}

func (ud *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

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

func (ud *userUseCase) LoginUserCase(authData user.LoginModel) (token, name string, err error) {
	token, name, err = ud.userData.LoginUserData(authData)
	return token, name, err
}
