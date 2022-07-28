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
	resultGet, _ := pd.postingData.GetUser(result.ID)
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
// func (pu *postingUseCase) DeleteCase(postingID int) (row int, err error) {
// 	row, err = pu.postingData.DeleteData(postingID)
// 	return row, err
// }

func (pu *postingUseCase) DeleteCase(idPosting, idFromToken int) (row int, err error) {
	row, err = pu.postingData.DeleteDataById(idPosting, idFromToken)
	return row, err
}

func (pu *postingUseCase) GetPostingById(id int) (data domain.Posting, err error) {
	data, err = pu.postingData.SelectDataById(id)
	return data, err
}
func (pu *postingUseCase) UpdateData(data domain.Posting, idPosting, idFromToken int) (row int, err error) {
	reqData := map[string]interface{}{}
	if data.Content != "" {
		reqData["content"] = data.Content
	}
	if data.Image != "" {
		reqData["image"] = data.Image
	}
	row, err = pu.postingData.UpdateData(reqData, idPosting, idFromToken)
	return row, err
}
