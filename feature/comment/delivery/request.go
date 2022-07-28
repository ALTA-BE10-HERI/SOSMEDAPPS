package delivery

import (
	"cleanarch/domain"

	"gorm.io/gorm"
)

type CommentInsertFormat struct {
	gorm.Model
	Comment    string `json:"comment" form:"comment"`
	ID_Users   int
	ID_Posting int
}

func (ci *CommentInsertFormat) ToModel() domain.Comment {
	return domain.Comment{
		Comment:  ci.Comment,
		ID_Users: ci.ID_Users,
		Posting: domain.PostingComment{
			ID: ci.ID_Posting,
		},
	}
}
