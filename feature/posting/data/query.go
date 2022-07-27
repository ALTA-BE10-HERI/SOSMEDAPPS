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

func (pd *postingData) InsertData(newPosting domain.Posting) (row int, err error) {
	posting := FromDomain(newPosting)
	res := pd.db.Create(&posting)
	if res.Error == nil {
		return 0, res.Error
	}
	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected != 1 {
		return 0, errors.New("failed to insert data")
	}
	return int(res.RowsAffected), err
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
