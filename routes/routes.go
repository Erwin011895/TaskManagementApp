package routes

import (
	"github.com/Erwin011895/TaskManagementApp/handler"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, h handler.Handler) {
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.GetUsers)
	r.GET("/users/:id", h.GetUser)
	r.PUT("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUser)

	r.POST("/tasks", AuthRequired(h), h.CreateTask)
	r.GET("/tasks", h.GetTasks)
	r.GET("/tasks/:id", AuthRequired(h), h.GetTask)
	r.PUT("/tasks/:id", AuthRequired(h), h.UpdateTask)
	r.DELETE("/tasks/:id", AuthRequired(h), h.DeleteTask)

}
