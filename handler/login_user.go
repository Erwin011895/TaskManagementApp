package handler

import (
	"net/http"

	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) LoginUser(c *gin.Context) {
	reqBody := model.User{}
	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: err.Error(),
		})
	}

	users, err := h.repo.GetUsers(c, &repo.GetUsersParam{
		Email: reqBody.Email,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	t, err := CreateToken(&h.config, *users[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, CreateUserResponse{
		Token: t,
	})
}
