package usecase

import (
	"cleanarch/domain"
	"errors"

	"github.com/go-playground/validator/v10"
)

type commentUseCase struct {
	data      domain.CommentData
	validator *validator.Validate
}

func New(model domain.CommentData) domain.CommentUseCase {
	return &commentUseCase{
		data: model,
	}
}

func (cs *commentUseCase) AddComment(IDLogin int, newText domain.Comment) (domain.Comment, error) {
	if IDLogin == -1 {
		return domain.Comment{}, errors.New("invalid user")
	}

	res := cs.data.Insert(newText)
	if res.ID == 0 {
		return domain.Comment{}, errors.New("failed to create comment")
	}

	return res, nil

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
