package handler

import (
	"net/http"
	"strconv"

	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) UpdateTask(c *gin.Context) {
	reqBody := model.Task{}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: err.Error(),
		})
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64) // Convert id from string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: "Invalid id",
		})
	}

	reqBody.Id = id

	jwtClaim := GetClaim(c)
	users, err := h.repo.GetUsers(c, &repo.GetUsersParam{
		Email: jwtClaim.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: err.Error(),
		})
		return
	}

	user := users[0]

	tasks, err := h.repo.GetTasks(c, &repo.GetTasksParam{
		TaskId: reqBody.Id,
		Limit:  1,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
		return
	}

	if tasks[0].UserId != user.Id {
		c.JSON(http.StatusUnauthorized, pkg.Response{
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	// Query update using the reqBody variable
	_, err = h.repo.UpdateTask(c, &reqBody)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}
	c.JSON(http.StatusOK, successResponse)
}
