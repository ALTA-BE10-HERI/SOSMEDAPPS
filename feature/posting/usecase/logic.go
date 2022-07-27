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

func (pd *postingUseCase) AddPosting(data domain.Posting) (result domain.Posting, err error) {
	result, err = pd.postingData.InsertData(data)
	resultGet, _ := pd.postingData.GetDetailPosting(result.ID)
	return resultGet, err
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
