package data

import (
	"cleanarch/domain"
	"errors"
	"fmt"

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

func (cd *commentData) InsertData(input domain.Comment) (row int, err error) {
	comment := FromDomain(input)
	fmt.Println("comment ", comment)
	res := cd.db.Create(&comment)
	fmt.Println("result error create: ", res.Error)

	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected != 1 {
		return 0, errors.New("failet to create comment")
	}

	return int(res.RowsAffected), nil
}

func (cd *commentData) SelectCommentByIdPosting(idPosting, limitint, offsetint int) (data []domain.Comment, err error) {
	comment := []Comment{}
	res := cd.db.Limit(limitint).Offset(offsetint).Preload("User").Order("created_at DESC").Where("posting_id = ?", idPosting).Find(&comment)
	if res.Error != nil {
		return []domain.Comment{}, nil
	}
	return parseToArrComment(comment), nil
}
