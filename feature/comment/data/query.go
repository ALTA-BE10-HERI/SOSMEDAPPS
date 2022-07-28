package data

import (
	"cleanarch/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CommentData {
	return &commentData{
		db: db,
	}
}

func (cd *commentData) Insert(newText domain.Comment) (result domain.Comment, err error) {
	comment := FromDomain(newText)
	res := cd.db.Create(&comment)
	if res.Error != nil {
		return domain.Comment{}, res.Error
	}
	if res.RowsAffected != 1 {
		return domain.Comment{}, errors.New("failed to insert data")
	}
	return comment.ToDomain(), err
}

func (cd *commentData) GetComment() []domain.Comment {
	var data []Comment
	err := cd.db.Limit(10).Find(&data).Error

	if err != nil {
		log.Println("Cannot read comment", err.Error())
		return nil
	}

	return ParseToArrComment(data)
}

func (cd *commentData) Delete(IDComment int) (row int, err error) {
	res := cd.db.Delete(&Comment{}, IDComment)
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
