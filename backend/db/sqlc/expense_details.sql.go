// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: expense_details.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExpenseDetail = `-- name: CreateExpenseDetail :one
INSERT INTO expense_details (
    expense_id,
    member_id,
    actual_amount,
    shared_amount
) VALUES (
$1, $2, $3 ,$4 
) RETURNING id, expense_id, member_id, actual_amount, shared_amount, created_at, updated_at
`

type CreateExpenseDetailParams struct {
	ExpenseID    pgtype.Int8    `json:"expense_id"`
	MemberID     pgtype.Int8    `json:"member_id"`
	ActualAmount pgtype.Numeric `json:"actual_amount"`
	SharedAmount pgtype.Numeric `json:"shared_amount"`
}

func (q *Queries) CreateExpenseDetail(ctx context.Context, arg CreateExpenseDetailParams) (ExpenseDetail, error) {
	row := q.db.QueryRow(ctx, createExpenseDetail,
		arg.ExpenseID,
		arg.MemberID,
		arg.ActualAmount,
		arg.SharedAmount,
	)
	var i ExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.ExpenseID,
		&i.MemberID,
		&i.ActualAmount,
		&i.SharedAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getExpenseDetail = `-- name: GetExpenseDetail :one
SELECT id, expense_id, member_id, actual_amount, shared_amount, created_at, updated_at FROM expense_details WHERE expense_id = $1
`

func (q *Queries) GetExpenseDetail(ctx context.Context, expenseID pgtype.Int8) (ExpenseDetail, error) {
	row := q.db.QueryRow(ctx, getExpenseDetail, expenseID)
	var i ExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.ExpenseID,
		&i.MemberID,
		&i.ActualAmount,
		&i.SharedAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateExpenseDetail = `-- name: UpdateExpenseDetail :one
UPDATE expense_details 
SET actual_amount = $3, 
shared_amount = $4 
WHERE expense_id = $1 
AND member_id = $2
RETURNING id, expense_id, member_id, actual_amount, shared_amount, created_at, updated_at
`

type UpdateExpenseDetailParams struct {
	ExpenseID    pgtype.Int8    `json:"expense_id"`
	MemberID     pgtype.Int8    `json:"member_id"`
	ActualAmount pgtype.Numeric `json:"actual_amount"`
	SharedAmount pgtype.Numeric `json:"shared_amount"`
}

func (q *Queries) UpdateExpenseDetail(ctx context.Context, arg UpdateExpenseDetailParams) (ExpenseDetail, error) {
	row := q.db.QueryRow(ctx, updateExpenseDetail,
		arg.ExpenseID,
		arg.MemberID,
		arg.ActualAmount,
		arg.SharedAmount,
	)
	var i ExpenseDetail
	err := row.Scan(
		&i.ID,
		&i.ExpenseID,
		&i.MemberID,
		&i.ActualAmount,
		&i.SharedAmount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
