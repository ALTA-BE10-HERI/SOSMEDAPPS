package usecase

import (
	"cleanarch/domain"
	"errors"
)

type commentUseCase struct {
	data domain.CommentData
}

func New(dataComment domain.CommentData) domain.CommentUseCase {
	return &commentUseCase{
		data: dataComment,
	}
}

func (cs *commentUseCase) CreateData(input domain.Comment) (row int, err error) {
	if input.Comment == "" || input.PostingID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = cs.data.InsertData(input)
	return row, err
}

func (cs *commentUseCase) GetCommentByIdPosting(idPosting, limitint, offsetint int) (data []domain.Comment, err error) {
	data, err = cs.data.SelectCommentByIdPosting(idPosting, limitint, offsetint)
	return data, err
}
