package main

import (
	"database/sql"
	"log"

	"github.com/RoseNeezar/go-e-wallet/api"
	db "github.com/RoseNeezar/go-e-wallet/db/sqlc"
	"github.com/RoseNeezar/go-e-wallet/util"
	_ "github.com/lib/pq"
)


func main() {
 
	config, err := util.LoadConfig(".")	

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
