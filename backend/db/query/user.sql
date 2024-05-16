-- name: CreateUser :one
INSERT INTO users (
username, 
password, 
email 
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
  username = COALESCE(sqlc.narg(username), username),
  password = COALESCE(sqlc.narg(password), password),
  email = COALESCE(sqlc.narg(email), email),
  updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
WHERE 
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;