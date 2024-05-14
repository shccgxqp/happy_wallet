package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/shccgxqp/happt_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeam(t *testing.T,userID int64) Team {
	arg := CreateTeamParams{
		Owner: userID,
		TeamName : util.RandomString(5),
		Currency : util.RandomCurrency(),
		TeamMembers : util.RandomTeamMembers(3),
	}

	team, err := testQueries.CreateTeam(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, team)

	return team
}


func TestGetTeam(t *testing.T) {
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)
	team2, err := testQueries.GetTeam(context.Background(),team1.ID)
	require.NoError(t, err)
	require.Equal(t, team1, team2)
}

func TestUpdateTeam(t *testing.T) {
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)

	arg := UpdateTeamParams{
		ID: team1.ID,
		TeamName : util.RandomString(5),
		Currency : util.RandomCurrency(),
		TeamMembers : util.RandomTeamMembers(3),
	}

	team2, err := testQueries.UpdateTeam(context.Background(),arg)
  require.NoError(t, err)
	require.NotEmpty(t,team2 )

	require.Equal(t, team1.ID, team2.ID)
	require.Equal(t, arg.TeamName, team2.TeamName)
	require.Equal(t, arg.Currency, team2.Currency)

	expectedMembers, err := json.Marshal(arg.TeamMembers)
	require.NoError(t, err)
	actualMembers, err := json.Marshal(team2.TeamMembers)
	require.NoError(t, err)
	require.JSONEq(t, string(expectedMembers), string(actualMembers))

	require.WithinDuration(t, team1.CreatedAt, team2.CreatedAt,time.Second)
	require.WithinDuration(t, team1.UpdatedAt, team2.UpdatedAt,time.Second)
}

func TestDeleteTeam(t *testing.T){
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)
	err := testQueries.DeleteTeam(context.Background(),team1.ID)
	require.NoError(t, err)

	team2, err := testQueries.GetTeam(context.Background(),team1.ID)
	require.Error(t, err)
	require.Empty(t, team2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListTeam(t *testing.T) {
	for i := 0; i < 6; i++ {
		user := createRandomUser(t)
		createRandomTeam(t, user.ID)
	}

	arg := ListTeamsParams{
		Limit: 5,
		Offset: 5,
	}

	teams, err := testQueries.ListTeams(context.Background(),arg)
	require.NoError(t, err)
	require.Len(t, teams, 5)

	for _,team := range teams {
		require.NotEmpty(t, team)
		require.NotEmpty(t, team.ID)
		require.NotEmpty(t, team.TeamName)
		require.NotEmpty(t, team.Currency)
		require.NotEmpty(t, team.TeamMembers)
		require.NotEmpty(t, team.CreatedAt)
		require.NotEmpty(t, team.UpdatedAt)
	}
}