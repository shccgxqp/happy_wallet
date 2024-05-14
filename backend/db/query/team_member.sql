-- name: CreateTeamMember :one
INSERT INTO team_members (
    team_id,
    member_name,
    linked_user_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTeamMembers :many
SELECT * FROM team_members
WHERE team_id = $1;

-- name: GetTeamMemberByID :one
SELECT * FROM team_members
WHERE id = $1;

-- name: UpdateTeamMember :one
UPDATE team_members
    SET linked_user_id = $3,
    member_name = $4
WHERE id = $1
AND team_id = $2
RETURNING *;