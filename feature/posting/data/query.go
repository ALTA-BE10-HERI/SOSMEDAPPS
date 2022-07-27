package data

import (
	"cleanarch/domain"
	"errors"
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

func (pd *postingData) GetDetailPosting(idPosting int) (result domain.Posting, err error) {
	var tmp Posting
	res := pd.db.Preload("User").Where("id = ?", idPosting).First(&tmp)
	if res.Error != nil {
		return domain.Posting{}, res.Error
	}
	return tmp.ToDomain(), nil
}

func (pd *postingData) InsertData(newPosting domain.Posting) (result domain.Posting, err error) {
	posting := FromDomain(newPosting)
	res := pd.db.Create(&posting)
	if res.Error != nil {
		return domain.Posting{}, res.Error
	}
	if res.RowsAffected != 1 {
		return domain.Posting{}, errors.New("failed to insert data")
	}
	return posting.ToDomain(), err
}

func (pd *postingData) GetPosting() []domain.Posting {
	var tmp []Posting
	err := pd.db.Limit(10).Find(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return nil
	}

	return ParseToArrPosting(tmp)
}

func (pd *postingData) DeleteData(postingID int) (row int, err error) {
	res := pd.db.Delete(&Posting{}, postingID)
	if res.Error != nil {
		log.Println("cannot delete data", res.Error.Error())
		return 0, res.Error
	}
	if res.RowsAffected < 1 {
		log.Println("no data deleted", res.Error.Error())
		return 0, errors.New("dailed to data deleted")
	}
	return int(res.RowsAffected), nil
}
