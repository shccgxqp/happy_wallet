package db

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeamMembers(t *testing.T, team Team) []TeamMember {
	var members []string
	json.Unmarshal(team.TeamMembers, &members)
	teamMembers := make([]TeamMember, 0, len(members))

	for _, member := range members {
		arg := CreateTeamMemberParams{
			TeamID:       pgtype.Int8{Int64: team.ID, Valid: true},
			MemberName:   string(member),
			LinkedUserID: pgtype.Int8{},
		}

		TeamMember, err := testStore.CreateTeamMember(context.Background(), arg)
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

	teamID := pgtype.Int8{Int64: team.ID, Valid: true}
	members, err := testStore.GetTeamMembers(context.Background(), teamID)
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
			TeamID:     pgtype.Int8{Int64: team.ID, Valid: true},
			MemberName: util.RandomUsername(),
		}

		TeamMember, err := testStore.UpdateTeamMember(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, TeamMember)
		require.Equal(t, arg.MemberName, TeamMember.MemberName)
		require.Equal(t, arg.TeamID, TeamMember.TeamID)
	}
}
