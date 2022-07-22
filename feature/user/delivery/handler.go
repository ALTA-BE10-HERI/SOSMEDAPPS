package delivery

import (
	"cleanarch/domain"
	"cleanarch/feature/common"
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

	e.POST("/user", handler.InsertUser())
	e.GET("/user", handler.GetAllUser())
	e.GET("/profile", handler.GetProfile(), _middleware.JWTMiddleware())
	e.POST("/login", handler.LoginAuth())

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
func (uh *userHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp, err := uh.userUsecase.GetAll()

		if err != nil {
			log.Println("cannot get all data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		if tmp == nil {
			return c.JSON(http.StatusInternalServerError, "error from database")
		}

		res := map[string]interface{}{
			"message": "succes get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}
func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
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
