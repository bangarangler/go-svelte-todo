package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	// "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/bangarangler/go-svelte-todo/pg"
)

type Resolver struct {
	Repository pg.Repository
}

func (r *mutationResolver) CreateUser(ctx context.Context, data UserInput) (*pg.User, error) {
	user, err := r.Repository.CreateUser(ctx, pg.CreateUserParams{
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, data UserInput) (*pg.User, error) {
	user, err := r.Repository.UpdateUser(ctx, pg.UpdateUserParams{
		ID:    id,
		Name:  data.Name,
		Email: data.Email,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*pg.User, error) {
	user, err := r.Repository.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, data TodoInput) (*pg.Todo, error) {
	todo, err := r.Repository.CreateTodo(ctx, pg.CreateTodoParams{
		Title:       data.Title,
		Description: pg.StringPtrToNullString(data.Description),
	})
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id int64, data TodoInput) (*pg.Todo, error) {
	todo, err := r.Repository.UpdateTodo(ctx, pg.UpdateTodoParams{
		Title:       data.Title,
		Description: pg.StringPtrToNullString(data.Description),
	})
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int64) (*pg.Todo, error) {
	todo, err := r.Repository.DeleteTodo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *queryResolver) User(ctx context.Context, id int64) (*pg.User, error) {
	user, err := r.Repository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]pg.User, error) {
	return r.Repository.ListUsers(ctx)
}

func (r *queryResolver) Todo(ctx context.Context, id int64) (*pg.Todo, error) {
	todo, err := r.Repository.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]pg.Todo, error) {
	return r.Repository.ListTodos(ctx)
}

func (r *todoResolver) Description(ctx context.Context, obj *pg.Todo) (*string, error) {
	var d string
	if obj.Description.Valid {
		d = obj.Description.String
		return &d, nil
	}
	return nil, nil
}

// func (r *__DirectiveResolver) IsRepeatable(ctx context.Context, obj *introspection.Directive) (bool, error) {
// 	panic("not implemented")
// }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// __Directive returns __DirectiveResolver implementation.
// func (r *Resolver) __Directive() __DirectiveResolver { return &__DirectiveResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// type __DirectiveResolver struct{ *Resolver }
