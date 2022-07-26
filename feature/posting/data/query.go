package data

import (
	"cleanarch/domain"
	"log"

	"gorm.io/gorm"
)

type postingData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.PostingData {
	return &postingData{
		db: db,
	}
}

func (pd *postingData) Insert(newPosting domain.Posting) domain.Posting {
	var cnv = FromDomain(newPosting)
	err := pd.db.Create(&cnv).Error
	if err != nil {
		log.Println("cannot create object", err.Error())
		return domain.Posting{}
	}

	return cnv.ToDomain()
}

func (pd *postingData) GetPosting() []domain.Posting {
	var tmp []Posting
	err := pd.db.Find(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return nil
	}

	return ParseToArrPosting(tmp)
}
