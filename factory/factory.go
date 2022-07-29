package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "cleanarch/feature/user/data"
	userDelivery "cleanarch/feature/user/delivery"
	us "cleanarch/feature/user/usecase"

	pd "cleanarch/feature/posting/data"
	postingDelivery "cleanarch/feature/posting/delivery"
	pu "cleanarch/feature/posting/usecase"

	cd "cleanarch/feature/comment/data"
	commentDelivery "cleanarch/feature/comment/delivery"
	cu "cleanarch/feature/comment/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.UserLogic(userData, validator)
	userDelivery.New(e, useCase)

	postingData := pd.New(db)
	PostingUserCase := pu.New(postingData)
	postingDelivery.New(e, PostingUserCase)

	commentData := cd.New(db)
	CommentUseCase := cu.New(commentData)
	commentDelivery.New(e, CommentUseCase)
}
