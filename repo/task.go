package repo

import (
	"context"
	"strings"

	"github.com/Erwin011895/TaskManagementApp/model"
)

type GetTasksParam struct {
	TaskId       int64
	UserId       int64
	OrderDueDate string
	Limit        int
	Offset       int
}

func (r PostgresRepo) GetTasks(ctx context.Context, param *GetTasksParam) (result []*model.Task, err error) {
	query := "SELECT * FROM tasks"

	// construct where queries
	whereQueries := []string{}
	queryArgs := []interface{}{}
	if param.TaskId > 0 {
		queryArgs = append(queryArgs, param.TaskId)
		whereQueries = append(whereQueries, "(id = ?)")
	}
	if param.UserId > 0 {
		whereQueries = append(whereQueries, "(user_id = ?)")
		queryArgs = append(queryArgs, param.UserId)
	}

	// append all where queries
	if len(whereQueries) > 0 {
		query += " WHERE " + strings.Join(whereQueries, " AND ")
	}

	// order by due date
	if len(param.OrderDueDate) > 0 {
		if param.OrderDueDate == "ASC" || param.OrderDueDate == "DESC" {
			query += " ORDER BY due_date " + param.OrderDueDate
		}
	}

	// limit & offset
	if param.Limit > 0 {
		query += " LIMIT ?"
		queryArgs = append(queryArgs, param.Limit)
	}
	if param.Offset > 0 {
		query += " OFFSET ?"
		queryArgs = append(queryArgs, param.Offset)
	}

	query = r.DB.Rebind(query)
	rows, err := r.DB.Queryx(query, queryArgs...)
	if err != nil {
		return
	}

	result = []*model.Task{}
	for rows.Next() {
		place := model.Task{}
		rows.StructScan(&place)
		result = append(result, &place)
	}
	return
}

func (r PostgresRepo) CreateTask(ctx context.Context, task *model.Task) (err error) {
	if task == nil {
		return
	}

	// construct query
	query := "INSERT INTO tasks (user_id, title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	queryArgs := []interface{}{task.UserId, task.Title, task.Description, task.Status}

	// db transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	query = r.DB.Rebind(query)
	_, err = tx.Exec(query, queryArgs...)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (r PostgresRepo) UpdateTask(ctx context.Context, task *model.Task) (result *model.Task, err error) {
	if task == nil {
		return
	}

	if task.Id == 0 {
		return
	}

	// construct query
	query := "UPDATE tasks SET title = ?, description = ?, status = ?, updated_at = NOW(), deleted_at = ? WHERE (id = ?)"
	queryArgs := []interface{}{task.Title, task.Description, task.Status, task.DeletedAt, task.Id}

	// db transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	query = r.DB.Rebind(query)
	_, err = tx.Exec(query, queryArgs...)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return task, nil
}
