-- name: CreateTeam :one
INSERT INTO teams (
owner,
team_name, 
currency,
team_members
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateTeam :one
UPDATE teams
  set team_name = $2,
  currency = $3,
  team_members = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM teams WHERE id = $1;