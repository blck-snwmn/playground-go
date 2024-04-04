-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
  id, name, bio
) VALUES (
  ?, ?, ?
);