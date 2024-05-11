package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) DeleteUser(c *gin.Context) {
	reqBody := &model.User{}

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

	users, err := h.repo.GetUsers(c, &repo.GetUsersParam{
		UserId: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, pkg.Response{
			Message: "User not found",
		})
	}
	user := users[0]

	jwtClaim := GetClaim(c)
	if jwtClaim.Email != user.Email {
		c.JSON(http.StatusUnauthorized, pkg.Response{
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	now := time.Now()
	user.DeletedAt = &now

	// Query update using the reqBody variable
	_, err = h.repo.UpdateUser(c, user)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}
	c.JSON(http.StatusOK, successResponse)
}
