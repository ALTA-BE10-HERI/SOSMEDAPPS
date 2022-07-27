package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	_helper "cleanarch/helper"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type postingHandler struct {
	postingUsercase domain.PostingUseCase
}

func New(e *echo.Echo, ps domain.PostingUseCase) {
	handler := &postingHandler{
		postingUsercase: ps,
	}

	e.POST("post", handler.InsertPosting(), _middleware.JWTMiddleware())
	e.GET("/post", handler.GetAllPosting())
	e.DELETE("/post/:id", handler.DeleteData(), _middleware.JWTMiddleware())
}

func (ph *postingHandler) InsertPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		idFromToken, _ := _middleware.ExtractData(c)
		err := c.Bind(&tmp)
		tmp.ID_Users = idFromToken

		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data"))
		}

		dataPosting := tmp.ToModel()
		result, errCreate := ph.postingUsercase.AddPosting(dataPosting)
		if errCreate != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to create posting"))
		}
		fmt.Println(result)
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModel(result)))

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
