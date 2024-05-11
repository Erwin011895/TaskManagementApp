package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/Erwin011895/TaskManagementApp/model"
)

type GetUsersParam struct {
	Cols   []string
	Limit  int
	Offset int
	UserId int64
	Email  string
}

func (r PostgresRepo) GetUsers(ctx context.Context, param *GetUsersParam) (result []*model.User, err error) {
	query := "SELECT * FROM users"

	if len(param.Cols) > 0 {
		query = fmt.Sprintf("SELECT %s FROM users", strings.Join(param.Cols, ", "))
	}

	// construct where queries
	whereQueries := []string{}
	queryArgs := []interface{}{}
	if param.UserId > 0 {
		queryArgs = append(queryArgs, param.UserId)
		whereQueries = append(whereQueries, "(id = ?)")
	}

	// append all where queries
	if len(whereQueries) > 0 {
		query += " WHERE " + strings.Join(whereQueries, " AND ")
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

	result = []*model.User{}
	for rows.Next() {
		place := model.User{}
		rows.StructScan(&place)
		result = append(result, &place)
	}
	return
}

func (r PostgresRepo) CreateUser(ctx context.Context, user *model.User) (err error) {
	if user == nil {
		return
	}

	// construct query
	query := "INSERT INTO users (email, name, password, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())"
	queryArgs := []interface{}{user.Email, user.Name, user.Password}

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

func (r PostgresRepo) UpdateUser(ctx context.Context, user *model.User) (result *model.User, err error) {
	if user == nil {
		return
	}

	if user.Id == 0 {
		return
	}

	// construct query
	query := "UPDATE users SET password = ?, name = ?, updated_at = NOW(), deleted_at = ? WHERE (id = ?)"
	queryArgs := []interface{}{user.Password, user.Name, user.DeletedAt, user.Id}

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

	return user, nil
}
