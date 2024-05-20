package db

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeam(t *testing.T, userID int64) Team {
	arg := CreateTeamParams{
		Owner:       userID,
		TeamName:    util.RandomString(5),
		Currency:    util.RandomCurrency(),
		TeamMembers: util.RandomTeamMembers(3),
	}

	team, err := testStore.CreateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team)

	return team
}

func TestGetTeam(t *testing.T) {
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)
	team2, err := testStore.GetTeam(context.Background(), team1.ID)
	require.NoError(t, err)
	require.Equal(t, team1, team2)
}

func TestUpdateTeam(t *testing.T) {
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)

	arg := UpdateTeamParams{
		ID:          team1.ID,
		TeamName:    util.RandomString(5),
		Currency:    util.RandomCurrency(),
		TeamMembers: util.RandomTeamMembers(3),
	}

	team2, err := testStore.UpdateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team2)

	require.Equal(t, team1.ID, team2.ID)
	require.Equal(t, arg.TeamName, team2.TeamName)
	require.Equal(t, arg.Currency, team2.Currency)

	var expectedMembers []string
	var actualMembers []string

	err = json.Unmarshal(arg.TeamMembers, &expectedMembers)
	require.NoError(t, err)
	err = json.Unmarshal(team2.TeamMembers, &actualMembers)
	require.NoError(t, err)

	require.ElementsMatch(t, expectedMembers, actualMembers)

	require.WithinDuration(t, team1.CreatedAt.Time, team2.CreatedAt.Time, time.Second)
	require.WithinDuration(t, team1.UpdatedAt.Time, team2.UpdatedAt.Time, time.Second)
}

func TestDeleteTeam(t *testing.T) {
	user := createRandomUser(t)
	team1 := createRandomTeam(t, user.ID)
	err := testStore.DeleteTeam(context.Background(), team1.ID)
	require.NoError(t, err)

	team2, err := testStore.GetTeam(context.Background(), team1.ID)
	require.Error(t, err)
	require.Empty(t, team2)
	require.EqualError(t, err, ErrRecordNotFound.Error())
}

func TestListTeam(t *testing.T) {
	for i := 0; i < 6; i++ {
		user := createRandomUser(t)
		createRandomTeam(t, user.ID)
	}

	arg := ListTeamsParams{
		Limit:  5,
		Offset: 5,
	}

	teams, err := testStore.ListTeams(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, teams, 5)

	for _, team := range teams {
		require.NotEmpty(t, team)
		require.NotEmpty(t, team.ID)
		require.NotEmpty(t, team.TeamName)
		require.NotEmpty(t, team.Currency)
		require.NotEmpty(t, team.TeamMembers)
		require.NotEmpty(t, team.CreatedAt)
		require.NotEmpty(t, team.UpdatedAt)
	}
}
