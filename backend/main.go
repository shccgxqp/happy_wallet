package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/shccgxqp/happt_wallet/backend/api"
	db "github.com/shccgxqp/happt_wallet/backend/db/sqlc"
	"github.com/shccgxqp/happt_wallet/backend/util"
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
	server := api.NewServer(store)
	
	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:",err)
	}
}
