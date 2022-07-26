package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	_helper "cleanarch/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentUsecase domain.CommentUseCase
}

func New(e *echo.Echo, cs domain.CommentUseCase) {
	handler := &commentHandler{
		commentUsecase: cs,
	}

	e.POST("/comments", handler.InsertComment(), _middleware.JWTMiddleware())
	e.GET("/comments", handler.GetAllComment())
	e.DELETE("/comments", handler.DeleteComment(), _middleware.JWTMiddleware())
}
func (ch *commentHandler) InsertComment() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp CommentInsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		id, _ := _middleware.ExtractData(c)
		data, err := ch.commentUsecase.AddComment(id, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot create comment", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Success create comment",
			"data":    data,
		})
	}
}

func (ch *commentHandler) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ch.commentUsecase.GetAllComment()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := _middleware.ExtractData(c)
		if token == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := ch.commentUsecase.DeleteComment(token)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
