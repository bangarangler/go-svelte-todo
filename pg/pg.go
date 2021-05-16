package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository interface {
	// user queries
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	// todo queires
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	DeleteTodo(ctx context.Context, id int64) (Todo, error)
	GetTodo(ctx context.Context, id int64) (Todo, error)
	ListTodos(ctx context.Context) ([]Todo, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
	ListTodosByUserID(ctx context.Context, userID int64) ([]Todo, error)
}

type repoSvc struct {
	*Queries
	db *sql.DB
}

func (r *repoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback: %v", err, rbErr)
		}
	} else {
		err = tx.Commit()
	}
	return err
}

func NewRepository(db *sql.DB) Repository {
	return &repoSvc{
		Queries: New(db),
		db:      db,
	}
}

func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

func StringPtrToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
