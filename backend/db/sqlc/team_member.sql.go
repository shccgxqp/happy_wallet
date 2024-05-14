// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: team_member.sql

package db

import (
	"context"
	"database/sql"
)

const createTeamMember = `-- name: CreateTeamMember :one
INSERT INTO team_members (
    team_id,
    member_name,
    linked_user_id
) VALUES (
  $1, $2, $3
) RETURNING id, team_id, member_name, linked_user_id, created_at, updated_at
`

type CreateTeamMemberParams struct {
	TeamID       sql.NullInt64 `json:"team_id"`
	MemberName   string        `json:"member_name"`
	LinkedUserID sql.NullInt64 `json:"linked_user_id"`
}

func (q *Queries) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, createTeamMember, arg.TeamID, arg.MemberName, arg.LinkedUserID)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.MemberName,
		&i.LinkedUserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeamMemberByID = `-- name: GetTeamMemberByID :one
SELECT id, team_id, member_name, linked_user_id, created_at, updated_at FROM team_members
WHERE id = $1
`

func (q *Queries) GetTeamMemberByID(ctx context.Context, id int64) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, getTeamMemberByID, id)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.MemberName,
		&i.LinkedUserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeamMembers = `-- name: GetTeamMembers :many
SELECT id, team_id, member_name, linked_user_id, created_at, updated_at FROM team_members
WHERE team_id = $1
`

func (q *Queries) GetTeamMembers(ctx context.Context, teamID sql.NullInt64) ([]TeamMember, error) {
	rows, err := q.db.QueryContext(ctx, getTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TeamMember{}
	for rows.Next() {
		var i TeamMember
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.MemberName,
			&i.LinkedUserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTeamMember = `-- name: UpdateTeamMember :one
UPDATE team_members
    SET linked_user_id = $3,
    member_name = $4
WHERE id = $1
AND team_id = $2
RETURNING id, team_id, member_name, linked_user_id, created_at, updated_at
`

type UpdateTeamMemberParams struct {
	ID           int64         `json:"id"`
	TeamID       sql.NullInt64 `json:"team_id"`
	LinkedUserID sql.NullInt64 `json:"linked_user_id"`
	MemberName   string        `json:"member_name"`
}

func (q *Queries) UpdateTeamMember(ctx context.Context, arg UpdateTeamMemberParams) (TeamMember, error) {
	row := q.db.QueryRowContext(ctx, updateTeamMember,
		arg.ID,
		arg.TeamID,
		arg.LinkedUserID,
		arg.MemberName,
	)
	var i TeamMember
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.MemberName,
		&i.LinkedUserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
