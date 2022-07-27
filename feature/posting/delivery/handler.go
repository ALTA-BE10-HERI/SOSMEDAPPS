package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	_helper "cleanarch/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type postingHandler struct {
	postingUsercase domain.PostingUserCase
}

func New(e *echo.Echo, ps domain.PostingUserCase) {
	handler := &postingHandler{
		postingUsercase: ps,
	}

	e.POST("user/posting", handler.InsertPosting(), _middleware.JWTMiddleware())
	e.GET("/homepage", handler.GetAllPosting())
}

func (ph *postingHandler) InsertPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := _middleware.ExtractData(c)
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := ph.postingUsercase.AddPosting(id, tmp.ToModel())

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

func (ph *postingHandler) GetAllPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ph.postingUsercase.GetAllPosting()

		if err != nil {
			log.Println("cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    data,
		})
	}
}

func (ph *postingHandler) DeleteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := _middleware.ExtractData(c)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("error read input"))
		}
		row, errDel := ph.postingUsercase.DeleteCase(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data"))
		}
		if row == 0 {
			return c.JSON(http.StatusNotFound, _helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success delete data"))
	}
}
