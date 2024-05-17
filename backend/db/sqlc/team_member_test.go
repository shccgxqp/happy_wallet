package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeamMembers(t *testing.T, team Team) []TeamMember {
	var members []string
	json.Unmarshal(team.TeamMembers.RawMessage, &members)
	teamMembers := make([]TeamMember, 0, len(members))

	for _, member := range members {
		arg := CreateTeamMemberParams{
			TeamID:       sql.NullInt64{Int64: team.ID, Valid: true},
			MemberName:   string(member),
			LinkedUserID: sql.NullInt64{},
		}

		TeamMember, err := testQueries.CreateTeamMember(context.Background(), arg)
		teamMembers = append(teamMembers, TeamMember)

		require.NoError(t, err)
		require.NotEmpty(t, TeamMember)
		require.Equal(t, team.ID, TeamMember.TeamID.Int64)
		require.Equal(t, string(member), TeamMember.MemberName)
		require.Equal(t, arg.LinkedUserID, TeamMember.LinkedUserID)
	}

	return teamMembers
}

func TestCreateTeamAndMembers(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	createRandomTeamMembers(t, team)
}

func TestGetTeamMembers(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	createRandomTeamMembers(t, team)

	teamID := sql.NullInt64{Int64: team.ID, Valid: true}
	members, err := testQueries.GetTeamMembers(context.Background(), teamID)
	require.NoError(t, err)
	require.NotEmpty(t, members)
}

func TestUpdateTeamMember(t *testing.T) {
	user := createRandomUser(t)
	team := createRandomTeam(t, user.ID)
	members := createRandomTeamMembers(t, team)

	for _, member := range members {
		arg := UpdateTeamMemberParams{
			ID:         member.ID,
			TeamID:     sql.NullInt64{Int64: team.ID, Valid: true},
			MemberName: util.RandomUsername(),
		}

		TeamMember, err := testQueries.UpdateTeamMember(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, TeamMember)
		require.Equal(t, arg.MemberName, TeamMember.MemberName)
		require.Equal(t, arg.TeamID, TeamMember.TeamID)
	}
}
