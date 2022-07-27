package usecase

import (
	"cleanarch/domain"
	"errors"
)

type postingUseCase struct {
	postingData domain.PostingData
}

func New(pd domain.PostingData) domain.PostingUserCase {
	return &postingUseCase{
		postingData: pd,
	}
}

func (pu *postingUseCase) AddPosting(userID int, newPosting domain.Posting) (domain.Posting, error) {
	if newPosting.ID_Users == 0 {
		return domain.Posting{}, errors.New("userID is empty")
	}
	return pu.postingData.Insert(newPosting), nil
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
