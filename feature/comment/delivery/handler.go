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
		idFromToken, _ := _middleware.ExtractData(c)
		err := c.Bind(&tmp)
		tmp.ID_Users = idFromToken
		tmp.ID_Posting = tmp.ToModel().ID_Posting

		if err != nil {
			c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data"))
		}

		dataComment := tmp.ToModel()
		result, errCreate := ch.commentUsecase.AddComment(dataComment)
		if errCreate != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to create comment"))
		}
		fmt.Println(result)
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModel(result)))

	}
}

func (ch *commentHandler) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ch.commentUsecase.GetAllComment()

		if err != nil {
			log.Println("cannot proccess data", err)
			c.JSON(http.StatusInternalServerError, err)
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
