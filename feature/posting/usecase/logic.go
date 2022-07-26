package usecase

import (
	"cleanarch/domain"
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

func (pu *postingUseCase) GetAllData(limit, offset int) (data []domain.Posting, err error) {
	res, err := pu.postingData.SelectData(limit, offset)
	return res, err
}
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
