package usecase

import (
	"cleanarch/domain"
	"errors"
)

type postingUseCase struct {
	postingData domain.PostingData
}

func New(pd domain.PostingData) domain.PostingUseCase {
	return &postingUseCase{
		postingData: pd,
	}
}

// func (pu *postingUseCase) AddPosting(userID int, newPosting domain.Posting) (domain.Posting, error) {
// 	if newPosting.ID_Users == 0 {
// 		return domain.Posting{}, errors.New("userID is empty")
// 	}
// 	return pu.postingData.Insert(newPosting), nil
// }

func (pd *postingUseCase) AddPosting(data domain.Posting) (row int, err error) {
	if data.Content == "" && data.Image == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = pd.postingData.InsertData(data)
	return row, err

}

func (pu *postingUseCase) GetAllPosting() ([]domain.Posting, error) {
	res := pu.postingData.GetPosting()

	if len(res) == 0 {
		return nil, errors.New("no data")
	}

	return res, nil
}

// make logic delete for this function
func (pu *postingUseCase) DeleteCase(postingID int) (row int, err error) {
	row, err = pu.postingData.DeleteData(postingID)
	return row, err
}
