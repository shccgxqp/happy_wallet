package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/token"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"github.com/shccgxqp/happy_wallet/backend/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TOKEN_SYMMETRIC_KEY:   util.RandomString(32),
		ACCESS_TOKEN_DURATION: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, userID int64, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(userID, duration)
	require.NoError(t, err)

	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}

	return metadata.NewIncomingContext(context.Background(), md)
}
