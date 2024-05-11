package handler

import (
	"net/http"
	"strconv"

	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: "Invalid id",
		})
	}

	users, err := h.repo.GetUsers(c, &repo.GetUsersParam{
		Cols:   []string{"id", "username", "name", "created_at", "updated_at"},
		UserId: userId,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, pkg.Response{
			Message: "user not found",
		})
	}
	c.JSON(http.StatusOK, users[0])
}
