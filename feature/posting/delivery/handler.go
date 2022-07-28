package delivery

import (
	"cleanarch/domain"
	_middleware "cleanarch/feature/common"
	_helper "cleanarch/helper"
	"fmt"
	"net/http"
	"strconv"

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
	e.GET("/post/:id", handler.GetById())
	e.PUT("/post/:id", handler.Update(), _middleware.JWTMiddleware())
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

// func (ph *postingHandler) GetAllPosting() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		data, err := ph.postingUsercase.GetAllPosting()

// 		if err != nil {
// 			log.Println("cannot proces data", err)
// 			c.JSON(http.StatusInternalServerError, err)
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "success get data",
// 			"data":    data,
// 		})
// 	}
// }

func (ph *postingHandler) GetAllPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)

		res, err := ph.postingUsercase.GetAllData(limitint, offsetint)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("canot to read all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(res)))

	}
}

func (ph *postingHandler) DeleteData() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idPost, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, err := ph.postingUsercase.DeleteCase(idPost, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have acces"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (ph *postingHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idPosting, _ := strconv.Atoi(id)
		res, err := ph.postingUsercase.GetPostingById(idPosting)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get detail posting"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success ", FromModel(res)))
	}

}

func (ph *postingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idPosting, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		content := c.FormValue("content")
		images := c.FormValue("image")
		postReq := InsertFormat{
			Content: content,
			Image:   images,
		}

		dataPost := postReq.ToModel()
		row, errUpd := ph.postingUsercase.UpdateData(dataPost, idPosting, idFromToken)
		if errUpd != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
