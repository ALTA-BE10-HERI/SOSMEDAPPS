package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentUseCase domain.CommentUseCase
}

func New(e *echo.Echo, cs domain.CommentUseCase) {
	handler := &commentHandler{
		commentUseCase: cs,
	}

	e.POST("/comment/create", handler.InsertComment(), _middleware.JWTMiddleware())
	e.GET("/comment", handler.InsertComment())

}
func (ch *commentHandler) InsertComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp CommentInsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := ch.commentUseCase.AddComment(tmp.ToDomain())

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
		data, err := ch.commentUseCase.GetAllComment()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}
