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
	err := pd.db.Limit(10).Find(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return nil
	}

	return ParseToArrPosting(tmp)
}

// make query delete for this func
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
