package db

import (
	"context"
	"math/big"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomExpense(t *testing.T, teamID pgtype.Int8) Expense {
	amount := util.RandomInt(1000, 10000)
	amountNumeric := pgtype.Numeric{
		Int:   big.NewInt(amount * 100), // 乘以100来模拟两位小数
		Exp:   -2,
		Valid: true,
	}

	arg := CreateExpenseParams{
		TeamID:        teamID,
		Goal:          util.RandomString(5),
		Amount:        amountNumeric,
		Currency:      util.RandomCurrency(),
		SharingMethod: util.RandomString(10),
	}

	expense, err := testStore.CreateExpense(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.Equal(t, arg.TeamID, expense.TeamID)
	require.Equal(t, arg.Goal, expense.Goal)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.Currency, expense.Currency)
	require.Equal(t, arg.SharingMethod, expense.SharingMethod)

	require.NotZero(t, expense.ID)
	require.NotZero(t, expense.CreatedAt)
	require.NotZero(t, expense.UpdatedAt)

	return expense
}

func TestCreateExpense(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	createRandomTeamMembers(t, team)
	createRandomExpense(t, pgtype.Int8{Int64: team.ID, Valid: true})
}

func TestGetExpense(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	createRandomTeamMembers(t, team)
	expense := createRandomExpense(t, pgtype.Int8{Int64: team.ID, Valid: true})
	gotExpense, err := testStore.GetExpense(context.Background(), expense.ID)
	require.NoError(t, err)
	require.Equal(t, expense, gotExpense)
}

func TestListExpenses(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	createRandomTeamMembers(t, team)
	for i := 0; i < 5; i++ {
		createRandomExpense(t, pgtype.Int8{Int64: team.ID, Valid: true})
	}

	expenses, err := testStore.ListExpenses(context.Background(), pgtype.Int8{Int64: team.ID, Valid: true})
	require.NoError(t, err)
	require.NotEmpty(t, expenses)
}
