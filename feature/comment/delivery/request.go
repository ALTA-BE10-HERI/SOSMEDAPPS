package delivery

import "cleanarch/domain"

type CommentInsertFormat struct {
	ID_Users   int
	ID_Posting int
	Comment    string `json:"comment" form:"comment"`
}

func (ci *CommentInsertFormat) ToDomain() domain.Comment {
	return domain.Comment{
		ID_Users:   ci.ID_Users,
		ID_Posting: ci.ID_Posting,
		Comment:    ci.Comment,
	}
}
