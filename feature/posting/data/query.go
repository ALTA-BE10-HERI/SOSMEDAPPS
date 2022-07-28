package data

import (
	"cleanarch/domain"
	"errors"

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

func (pd *postingData) SelectDataById(id int) (data domain.Posting, err error) {
	tmp := Posting{}
	res := pd.db.Preload("User").Find(&tmp, id)
	if res.Error != nil {
		return domain.Posting{}, res.Error
	}
	return tmp.ToDomain(), nil
}

func (pd *postingData) GetUser(idPosting int) (result domain.Posting, err error) {
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

// func (pd *postingData) GetPosting() []domain.Posting {
// 	var tmp []Posting
// 	err := pd.db.Limit(10).Find(&tmp).Error
// 	if err != nil {
// 		log.Println("There is a problem with data", err.Error())
// 		return nil
// 	}

// 	return ParseToArrPosting(tmp)
// }

func (pd *postingData) SelectData(limit, offset int) (data []domain.Posting, err error) {
	dataPosting := []Posting{}
	res := pd.db.Preload("User").Limit(limit).Offset(offset).Find(&dataPosting)
	if res.Error != nil {
		return []domain.Posting{}, nil
	}
	return ParseToArrPosting(dataPosting), nil
}

func (pd *postingData) DeleteDataById(idPosting, idFromToken int) (row int, err error) {
	dataPosting := Posting{}
	cekID := pd.db.First(&dataPosting, idPosting)
	if cekID.Error != nil {
		return 0, cekID.Error
	}
	if idFromToken != int(dataPosting.UserID) {
		return -1, errors.New("you don`t have acces")
	}

	res := pd.db.Delete(&Posting{}, idPosting)
	if res.Error != nil {
		return 0, res.Error
	}
	if res.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(res.RowsAffected), nil
}

func (pd *postingData) UpdateData(data map[string]interface{}, idPosting, idFromToken int) (row int, err error) {
	dataPosting := Posting{}
	cekID := pd.db.First(&dataPosting, "id  = ?", idPosting)

	if cekID.Error != nil {
		return 0, cekID.Error
	}
	if dataPosting.UserID != idFromToken {
		return -1, errors.New("you don`t have access")
	}
	res := pd.db.Model(&Posting{}).Where("id = ? ", idPosting).Updates(&data)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected != 1 {
		return 0, errors.New("failed update")
	}
	return int(res.RowsAffected), nil
}
