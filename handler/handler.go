package handler

import (
	"github.com/Erwin011895/TaskManagementApp/config"
	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Handler interface {
	CreateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	CreateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type handlerImpl struct {
	config config.Config
	repo   repo.Repository
}

var (
	successResponse = pkg.Response{
		Message: "Success",
	}

	serverErrorResponse = pkg.Response{
		Message: "Something went wrong!",
	}
)

func InitHandler(db *sqlx.DB, conf config.Config) Handler {
	return handlerImpl{
		config: conf,
		repo: repo.PostgresRepo{
			DB:     db,
			Config: conf,
		},
	}
}
