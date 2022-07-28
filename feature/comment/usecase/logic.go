package usecase

import (
	"cleanarch/domain"
	// "cleanarch/feature/posting/usecase/logic"
	"errors"
)

type commentUseCase struct {
	commentData domain.CommentData
}

func New(cd domain.CommentData) domain.CommentUseCase {
	return &commentUseCase{
		commentData: cd,
	}
}

func (pu *domain.PostingUseCase) AddPosting(postById domain.Posting) (result domain.Posting, err error) {
	result, err = pd.postingData.InsertData(postById)
	resultGet, _ := pd.postingData.GetDetailPosting(result.ID)
	return resultGet, err
}

func (cs *commentUseCase) AddComment(data domain.Comment) (result domain.Comment, err error) {

}

func (cs *commentUseCase) GetAllComment() ([]domain.Comment, error) {

	res := cs.data.GetComment()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (cs *commentUseCase) DeleteComment(IDComment int) (row int, err error) {
	row, err = cs.data.Delete(IDComment)
	return row, err
}
