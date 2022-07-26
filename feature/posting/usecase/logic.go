package usecase

import (
	"cleanarch/domain"
	"errors"
)

type postingUseCase struct {
	data domain.PostingData
}

func New(model domain.PostingData) domain.PostingUserCase {
	return &postingUseCase{
		data: model,
	}
}

func (pu *postingUseCase) AddPosting(userID int, newPosting domain.Posting) (domain.Posting, error) {
	if newPosting.ID_Users == 0 {
		return domain.Posting{}, errors.New("userID is empty")
	}
	return pu.data.Insert(newPosting), nil
}

func (pu *postingUseCase) GetAllPosting() ([]domain.Posting, error) {
	res := pu.data.GetPosting()

	if len(res) == 0 {
		return nil, errors.New("no data")
	}

	return res, nil
}
