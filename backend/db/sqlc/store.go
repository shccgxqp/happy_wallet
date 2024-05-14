package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := New(tx)
	err = fn(queries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

// type expensesTxParams struct {
// 	Name      string  `json:"name"`
// 	Amount    float64 `json:"amount"`
// 	UserID    int64   `json:"user_id"`
// 	GroupCode string  `json:"group_code"`
// 	ShareType string  `json:"share_type"`   // 新增分享類型
// 	ShareValue float64 `json:"share_value"` // 新增分享值
// }

// // TransferTxResult contains the result of the TransferTx function.
// // The created Transfer record.
// // The FromAccount after its balance is subtracted.
// // The ToAccount after its its balance is added.
// // The FromEntry of the account which records that money is moving out of the FromAccount.
// // And the ToEntry of the account which records that money is moving in to the ToAccount.
// type expensesTxResult struct {
// 	Expenses Expense `json:"expenses"`
// 	Share Share `json:"share"`
// }


// func (store *Store) expensesTx(ctx context.Context, arg expensesTxParams) (expensesTxResult, error) {
// 	var result expensesTxResult
// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.Expenses, err = q.CreateExpenses(ctx, CreateExpensesParams{
// 			Name:      arg.Name,
// 			Amount:    arg.Amount,
// 			UserID:    arg.UserID,
// 			GroupCode: arg.GroupCode,
// 		})

// 		if err != nil {
// 			return err
// 		}

// 		result.Share, err = q.CreateShare(ctx, CreateShareParams{
// 			ExpenseID: result.Expenses.ID,
// 			UserID:    arg.UserID,
// 			ShareType: arg.ShareType,
// 			ShareValue: arg.ShareValue,

// 		})

// 		if err != nil {
// 			return err
// 		}
// 		return nil

// 		// TODO: update accounts' balance
// 	})

// 	return result, err
// }