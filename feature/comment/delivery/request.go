package delivery

import "cleanarch/domain"

type CommentInsertFormat struct {
	PostingID int    `json:"posting_id" form:"posting_id"`
	Comment   string `json:"comment" form:"comment"`
	UserID    int    `json:"user_id" form:"user_id"`
	Nama      string `json:"nama" form:"nama"`
}

func (ci *CommentInsertFormat) ToDomain() domain.Comment {
	return domain.Comment{
		PostingID: ci.PostingID,
		Comment:   ci.Comment,
		UserID:    ci.UserID,
		Nama:      ci.Nama,
	}
}
