package usecase

import (
	"cleanarch/domain"
	"cleanarch/feature/user"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator"
	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func UserLogic(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	// return &userUseCase{
	// 	userData: ud,
	// }
	return &userUseCase{
		userData: ud,
		validate: v,
	}
}

func (uc *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	if newUser.Nama == "" || newUser.Email == "" || newUser.Password == "" {
		return domain.User{}, errors.New("please make sure all fields are filled in correctly")
	}
	hashed, err := _bcrypt.GenerateFromPassword([]byte(newUser.Password), _bcrypt.DefaultCost)

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

func (uc *userUseCase) UpdateCase(input domain.User, idFromToken int) (row int, err error) {
	userReq := map[string]interface{}{}
	if input.Nama != "" {
		userReq["nama"] = input.Nama
	}
	if input.Email != "" {
		userReq["email"] = input.Email
	}
	if input.Password != "" {
		passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		if errorHash != nil {
			fmt.Println("Error hash", errorHash.Error())
		}
		userReq["password"] = string(passwordHashed)
	}
	row, err = uc.userData.UpdateData(userReq, idFromToken)
	return row, err
}
