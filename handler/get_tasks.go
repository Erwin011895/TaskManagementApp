package handler

import (
	"net/http"
	"strconv"

	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) GetTasks(c *gin.Context) {
	var err error
	userIdParam := c.Query("user_id")
	offsetParam := c.Query("offset")
	limitParam := c.Query("limit")

	var userId int64
	if userIdParam == "" {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: "Invalid user_id param",
		})
	}

	userId, err = strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: "Invalid user_id param",
		})
	}

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

	tasks, err := h.repo.GetTasks(c, &repo.GetTasksParam{
		UserId:       userId,
		OrderDueDate: "ASC",
		Limit:        limit,
		Offset:       offset,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	c.JSON(http.StatusOK, tasks)
}
