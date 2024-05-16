package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/pb"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
			Username:       req.GetUsername(),
			Password:       hashedPassword,
			Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr,ok:=err.(*pq.Error);ok{
			switch pqErr.Code.Name(){
				case "unique_violation":
					return nil, status.Errorf(codes.AlreadyExists, "username or email already exists: %s",err)
			}
			return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
		}
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
