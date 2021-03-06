// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package pg

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING id, name, email
`

type CreateUserParams struct {
	Name  string
	Email string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, name, email
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email FROM users
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3
WHERE id = $1
RETURNING id, name, email
`

type UpdateUserParams struct {
	ID    int64
	Name  string
	Email string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Name, arg.Email)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}
