package data

import (
	"cleanarch/domain"
	"cleanarch/feature/common"
	"cleanarch/feature/user"
	"errors"
	"log"

	_bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) domain.User {
	var cnv = FromModel(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("cannot create object", err.Error())
		return domain.User{}
	}

	return cnv.ToModel()
}

func (ud *userData) GetAll() []domain.User {
	var tmp []User
	err := ud.db.Find(&tmp).Error

	if err != nil {
		log.Println("cannot retrive object", err.Error())
		return nil
	}
	return ParseToArr(tmp)
}
func (ud *userData) GetSpecific(userID int) (domain.User, error) {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}, err
	}

	return tmp.ToModel(), nil
}
func (ud *userData) LoginUserData(authData user.LoginModel) (token, name string, err error) {
	userData := User{}
	result := ud.db.Where("email = ?", authData.Email).First(&userData)

	if result.Error != nil {
		return "", "", result.Error
	}

	if result.RowsAffected != 1 {
		return "", "", errors.New("failed to login")
	}

	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))
	if errCrypt != nil {
		return "", "", errors.New("password incorrect")
	}
	token = common.GenerateToken(int(userData.ID))

	return token, userData.Nama, nil
}
