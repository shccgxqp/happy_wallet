package gapi

import (
	"context"
	"database/sql"

	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/pb"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, status.Errorf(codes.NotFound, "user not found")
			}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil{
		return nil , status.Errorf(codes.Unauthenticated, "incorrect password")
	}

	accessToken,accessPayload, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.ACCESS_TOKEN_DURATION,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.REFRESH_TOKEN_DURATION,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
	ID:           refreshPayload.ID,
	UserID:       user.ID,
	RefreshToken: refreshToken,
	UserAgent:    "",
	ClientIp:     "",
	IsBlocked:    false,
	ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User: convertUser(user),
		SessionId: session.ID.String(),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		AccessTokenExpiresTime: timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresTime: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil
}