// Code generated by sqlc. DO NOT EDIT.

package pg

import (
	"database/sql"
)

type Todo struct {
	ID          int64
	Title       string
	Description sql.NullString
	UserID      int64
}

type User struct {
	ID    int64
	Name  string
	Email string
}