package usecase

import (
	"cleanarch/domain"
	"errors"
)

type commentUseCase struct {
	commentData domain.CommentData
}

func (cs *commentUseCase) AddComment(IDLogin int, newText domain.Comment) (domain.Comment, error) {
	if IDLogin == -1 {
		return domain.Comment{}, errors.New("invalid user")
	}

	res := cs.commentData.Insert(newText)
	if res.ID == 0 {
		return domain.Comment{}, errors.New("failed to create comment")
	}

	return res, nil

}

func (cs *commentUseCase) GetComment() ([]domain.Comment, error) {
	res := cs.commentData.GetComment()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}
