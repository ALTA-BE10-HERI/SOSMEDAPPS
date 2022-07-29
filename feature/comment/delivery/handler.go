package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	_helper "cleanarch/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentUsecase domain.CommentUseCase
}

func New(e *echo.Echo, cs domain.CommentUseCase) {
	handler := &commentHandler{
		commentUsecase: cs,
	}

	e.POST("/comments/:id", handler.InsertComment(), _middleware.JWTMiddleware())
	e.GET("/comments/:id", handler.GetAllComment())
	e.DELETE("/comments/:id", handler.DeleteComment(), _middleware.JWTMiddleware())
}
func (ch *commentHandler) InsertComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idPosting, _ := strconv.Atoi(id)
		commentReq := CommentInsertFormat{}
		err := c.Bind(&commentReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
		}
		idFromToken, _ := _middleware.ExtractData(c)
		commentReq.UserID = idFromToken
		commentReq.PostingID = idPosting
		dataComment := commentReq.ToDomain()

		row, errCreate := ch.commentUsecase.CreateData(dataComment)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
		}

		if errCreate != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add comment"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
func (ch *commentHandler) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)
		idPosting, _ := strconv.Atoi(id)

		res, errGet := ch.commentUsecase.GetCommentByIdPosting(idPosting, limitint, offsetint)
		if errGet != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data"))
		}
		log.Println("cek ", res)
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(res)))
	}
}

func (ch *commentHandler) DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idComment, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, errDelelete := ch.commentUsecase.DeleteCommentById(idComment, idFromToken)
		if errDelelete != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data user"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
