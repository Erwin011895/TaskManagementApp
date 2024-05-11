package handler

import (
	"net/http"

	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserResponse struct {
	Token string `json:"token,omitempty"`
}

func (h handlerImpl) CreateUser(c *gin.Context) {
	reqBody := model.User{}
	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: err.Error(),
		})
	}

	// encrypt password
	bytes, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
		return
	}
	reqBody.Password = string(bytes)

	err = h.repo.CreateUser(c, &reqBody)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	t, err := CreateToken(&h.config, reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, CreateUserResponse{
		Token: t,
	})
}
