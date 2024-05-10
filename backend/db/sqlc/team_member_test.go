package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/shccgxqp/happt_wallet/backend/util"
	"github.com/stretchr/testify/require"
)

func TestGetTeamMembers(t *testing.T){
	team :=createRandomTeam(t)
	createRandomTeamMembers(t, team)

	teamID := sql.NullInt64{Int64: team.ID, Valid: true}
	members, err := testQueries.GetTeamMembers(context.Background(), teamID)
	require.NoError(t, err)
	require.NotEmpty(t, members)
}

func TestUpdateTeamMermber(t *testing.T){
	team :=createRandomTeam(t)
	members := createRandomTeamMembers(t, team)

	for _, member := range members {
		arg:= updateTeamMemberParams{
			ID : member.ID,
			TeamID : sql.NullInt64{Int64: team.ID, Valid: true},
			MemberName : util.RandomUsername(),
		}

		TeamMember, err := testQueries.updateTeamMember(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, TeamMember)
		require.Equal(t, arg.MemberName, TeamMember.MemberName)
		require.Equal(t, arg.TeamID, TeamMember.TeamID)
	}
}

