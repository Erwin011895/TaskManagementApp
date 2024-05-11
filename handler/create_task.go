package handler

import (
	"net/http"

	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) CreateTask(c *gin.Context) {
	reqBody := model.Task{}
	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: err.Error(),
		})
		return
	}

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

	reqBody.UserId = user.Id
	reqBody.Status = "pending"

	err = h.repo.CreateTask(c, &reqBody)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	c.JSON(http.StatusOK, successResponse)
}
