package data

import (
	"cleanarch/domain"
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

func (cd *commentData) Insert(newText domain.Comment) domain.Comment {
	var cnv = FromDomain(newText)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("cannot create comment", err.Error())
		return domain.Comment{}
	}

	return cnv.ToDomain()
}

func (cd *commentData) GetComment() []domain.Comment {
	var data []Comment
	err := cd.db.Find(&data)

	if err.Error != nil {
		log.Println("Cannot read comment", err.Error.Error())
		return nil
	}

	return ParseToArrComment(data)
}
