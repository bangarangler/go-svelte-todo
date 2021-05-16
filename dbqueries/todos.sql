-- name: GetTodo :one
SELECT * from todos
WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY title;

-- name: CreateTodo :one
INSERT INTO todos (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;

-- name: ListTodosByUserID :many
SELECT todos.* FROM todos, users
WHERE users.id = todos.user_id AND todos.user_id = $1;
