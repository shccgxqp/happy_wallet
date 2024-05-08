// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: expense.sql

package db

import (
	"context"
	"database/sql"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (
    team_id,
    goal,
    amount,
    currency,
    sharing_method
) VALUES (
$1, $2, $3 ,$4 ,$5
) RETURNING id, team_id, goal, amount, currency, sharing_method, created_at, updated_at
`

type CreateExpenseParams struct {
	TeamID        sql.NullInt64 `json:"team_id"`
	Goal          string        `json:"goal"`
	Amount        string        `json:"amount"`
	Currency      string        `json:"currency"`
	SharingMethod string        `json:"sharing_method"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.TeamID,
		arg.Goal,
		arg.Amount,
		arg.Currency,
		arg.SharingMethod,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Goal,
		&i.Amount,
		&i.Currency,
		&i.SharingMethod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getExpense = `-- name: GetExpense :one
SELECT id, team_id, goal, amount, currency, sharing_method, created_at, updated_at FROM expenses WHERE id = $1
`

func (q *Queries) GetExpense(ctx context.Context, id int64) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Goal,
		&i.Amount,
		&i.Currency,
		&i.SharingMethod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, team_id, goal, amount, currency, sharing_method, created_at, updated_at FROM expenses WHERE team_id = $1
`

func (q *Queries) ListExpenses(ctx context.Context, teamID sql.NullInt64) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Expense{}
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.Goal,
			&i.Amount,
			&i.Currency,
			&i.SharingMethod,
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

const updateExpense = `-- name: UpdateExpense :one
UPDATE expenses SET 
    goal = $2,
    amount = $3,
    currency = $4,
    sharing_method = $5
WHERE id = $1 RETURNING id, team_id, goal, amount, currency, sharing_method, created_at, updated_at
`

type UpdateExpenseParams struct {
	ID            int64  `json:"id"`
	Goal          string `json:"goal"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	SharingMethod string `json:"sharing_method"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, updateExpense,
		arg.ID,
		arg.Goal,
		arg.Amount,
		arg.Currency,
		arg.SharingMethod,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Goal,
		&i.Amount,
		&i.Currency,
		&i.SharingMethod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
