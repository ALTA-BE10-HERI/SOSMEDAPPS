package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	"cleanarch/feature/user"
	_helper "cleanarch/helper"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}

	e.POST("/register", handler.InsertUser())
	e.GET("/myprofile", handler.GetProfile(), _middleware.JWTMiddleware())
	e.POST("/login", handler.LoginAuth())
	e.DELETE("/user/delete", handler.DeleteById(), _middleware.JWTMiddleware())
}
func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})
	}
}
func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := _middleware.ExtractData(c)
		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}
func (uh *userHandler) LoginAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		authData := user.LoginModel{}
		c.Bind(&authData)
		token, name, e := uh.userUsecase.LoginUserCase(authData)
		if e != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
		}

		data := map[string]interface{}{
			"token": token,
			"name":  name,
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", data))
	}
}

func (uh *userHandler) DeleteById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := uh.userUsecase.DeleteCase(idFromToken)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
