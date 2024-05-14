package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/shccgxqp/happt_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomExpense(t *testing.T,teamID sql.NullInt64) Expense{
	arg :=CreateExpenseParams{
		TeamID: teamID,
		Goal: util.RandomString(5),
		Amount: strconv.FormatFloat(util.RandomFloat(100,1000), 'f', 2, 64), 
		Currency: util.RandomCurrency(),
		SharingMethod: util.RandomString(10),
	}

	expense, err := testQueries.CreateExpense(context.Background(), arg)
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


func TestCreateExpense(t *testing.T){
	user := createRandomUser(t)
  team :=createRandomTeam(t, user.ID)
	createRandomTeamMembers(t,team)
  createRandomExpense(t,sql.NullInt64{Int64:team.ID,Valid:true})
}

func TestGetExpense(t *testing.T){
	user := createRandomUser(t)
  team :=createRandomTeam(t, user.ID)
	createRandomTeamMembers(t,team)
	expense := createRandomExpense(t,sql.NullInt64{Int64:team.ID,Valid:true})
	gotExpense, err := testQueries.GetExpense(context.Background(), expense.ID)
	require.NoError(t, err)
	require.Equal(t, expense, gotExpense)
}

func TestListExpenses(t *testing.T){
	user := createRandomUser(t)
  team :=createRandomTeam(t, user.ID)
	createRandomTeamMembers(t,team)
	for i := 0; i < 5; i++ {
		createRandomExpense(t,sql.NullInt64{Int64:team.ID,Valid:true})
	}

	expenses, err := testQueries.ListExpenses(context.Background(), sql.NullInt64{Int64:team.ID,Valid:true})
	require.NoError(t, err)
	require.NotEmpty(t, expenses)
}

