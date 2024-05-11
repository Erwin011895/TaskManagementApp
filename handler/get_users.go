package handler

import (
	"net/http"
	"strconv"

	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) GetUsers(c *gin.Context) {
	var err error
	offsetParam := c.Query("offset")
	limitParam := c.Query("limit")

	offset := 0
	limit := 10

	if len(offsetParam) > 0 {
		offset, err = strconv.Atoi(offsetParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.Response{
				Message: "Invalid offset param",
			})
		}
	}

	if len(limitParam) > 0 {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.Response{
				Message: "Invalid limit param",
			})
		}
	}

	users, err := h.repo.GetUsers(c, &repo.GetUsersParam{
		Cols:   []string{"id", "username", "name", "created_at", "updated_at"},
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	c.JSON(http.StatusOK, users)
}
