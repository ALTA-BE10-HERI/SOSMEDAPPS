package factory

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "cleanarch/feature/user/data"
	userDelivery "cleanarch/feature/user/delivery"
	us "cleanarch/feature/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.UserLogic(userData, validator)
	userDelivery.New(e, useCase)
}
