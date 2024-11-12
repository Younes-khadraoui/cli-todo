-- name: NewTodo :exec
INSERT INTO todos (
  content
) values (
  ?
);

-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = ?
