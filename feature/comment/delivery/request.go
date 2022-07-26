package delivery

import "cleanarch/domain"

type CommentInsertFormat struct {
	ID      int    `json:"id"`
	Comment string `json:"comment" form:"comment"`
}

func (ci *CommentInsertFormat) ToDomain() domain.Comment {
	return domain.Comment{
		ID:      ci.ID,
		Comment: ci.Comment,
	}
}
