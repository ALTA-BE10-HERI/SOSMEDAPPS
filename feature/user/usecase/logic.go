package usecase

import (
	"cleanarch/domain"
	"cleanarch/feature/user"
	"cleanarch/feature/user/data"
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

func New(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		validate: v,
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

func (uc *userUseCase) UpdateCase(userID int, newUser domain.User) (domain.User, error) {
	var cnv = data.FromModel(newUser)
	err := uc.validate.Struct(cnv)
	if err != nil {
		log.Println("Validation errror : ", err.Error())
		return domain.User{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	updated, err := uc.userData.UpdateData(userID, newUser)

	if err != nil {
		log.Println("User Usecase", err.Error())
		return domain.User{}, err
	}

	if updated.ID == 0 {
		return domain.User{}, errors.New("cannot update data")
	}

	return updated, nil
}
