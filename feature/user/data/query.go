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
	token, _ = common.GenerateToken(int(userData.ID))

	return token, userData.Nama, nil
}

func (ud *userData) DeleteData(userID int) (row int, err error) {
	res := ud.db.Delete(&User{}, userID)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return 0, errors.New("failed to delete data ")
	}
	return int(res.RowsAffected), nil
}

func (ud *userData) UpdateData(oldUser domain.User, newUser domain.User) domain.User {
	var cnv = FromModel(newUser)
	res := ud.db.Model(&User{}).Where("ID = ?", oldUser.ID).Updates(&cnv)
	if res.Error != nil {
		log.Println("cannot update data", res.Error.Error())
		return domain.User{}
	}
	if res.RowsAffected < 1 {
		log.Println("no data updated", res.Error.Error())
		return domain.User{}
	}
	return cnv.ToModel()
}
