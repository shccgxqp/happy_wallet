package gapi

import (
	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		CreatedAt:         timestamppb.New(user.CreatedAt),
		UpdatedAt:				 timestamppb.New(user.UpdatedAt),
	}
}