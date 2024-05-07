-- name: QueryEGroup :one
INSERT INTO e_group (
code, 
name, 
currency 
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetEGroup :one
SELECT * FROM e_group
WHERE id = $1 LIMIT 1;

-- name: ListEGroup :many
SELECT * FROM e_group
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateEGroup :one
UPDATE e_group
  set name = $2,
  currency = $3
WHERE code = $1
RETURNING *;

-- name: DeleteEGroup :exec
DELETE FROM e_group WHERE id = $1;