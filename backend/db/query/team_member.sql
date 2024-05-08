-- name: CreateTeamMember :one
INSERT INTO team_members (
    team_id,
    member_name,
    linked_account_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTeamMembers :many
SELECT * FROM team_members
WHERE team_id = $1;

-- name: GetTeamMember :one
SELECT * FROM team_members
WHERE id = $1
AND team_id = $2;

-- name: updateTeamMember :one
UPDATE team_members
    SET linked_account_id = $3,
    member_name = $4
WHERE id = $1
AND team_id = $2
RETURNING *;