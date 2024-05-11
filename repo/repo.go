package repo

import (
	"context"

	"github.com/Erwin011895/TaskManagementApp/config"
	"github.com/Erwin011895/TaskManagementApp/model"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateTask(ctx context.Context, task *model.Task) (err error)
	CreateUser(ctx context.Context, user *model.User) (err error)
	GetTasks(ctx context.Context, param *GetTasksParam) (result []*model.Task, err error)
	GetUsers(ctx context.Context, param *GetUsersParam) (result []*model.User, err error)
	UpdateTask(ctx context.Context, task *model.Task) (result *model.Task, err error)
	UpdateUser(ctx context.Context, user *model.User) (result *model.User, err error)
}

type PostgresRepo struct {
	DB     *sqlx.DB
	Config config.Config
}
