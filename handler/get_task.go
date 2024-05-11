package handler

import (
	"net/http"
	"strconv"

	"github.com/Erwin011895/TaskManagementApp/pkg"
	"github.com/Erwin011895/TaskManagementApp/repo"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func (h handlerImpl) GetTask(c *gin.Context) {
	taskId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.Response{
			Message: "Invalid id",
		})
	}

	tasks, err := h.repo.GetTasks(c, &repo.GetTasksParam{
		TaskId: taskId,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, serverErrorResponse)
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, pkg.Response{
			Message: "task not found",
		})
	}
	c.JSON(http.StatusOK, tasks[0])
}
