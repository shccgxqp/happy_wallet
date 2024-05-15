package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/shccgxqp/happy_wallet/backend/api"
	db "github.com/shccgxqp/happy_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happy_wallet/backend/gapi"
	"github.com/shccgxqp/happy_wallet/backend/pb"
	"github.com/shccgxqp/happy_wallet/backend/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:",err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
		if err != nil {
			log.Fatal("cannot connect to db:",err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config,store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
			log.Fatal("cannot create grpc server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHappyWalletServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_SERVER_ADDRESS)
    if err != nil {
        log.Fatal("cannot listen to grpc server:", err)
    }

    log.Printf("start gRPC server at %s", listener.Addr().String())
    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatal("cannot serve grpc server:", err)
    }
}

func runGinServer(config util.Config,store db.Store) {
	server,err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:",err)
	}
	
	err = server.Start(config.HTTP_SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:",err)
	}
}