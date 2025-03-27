package main

import (
	"database/sql"
	"log"

	"github.com/IamDushu/Harbor-Health-Server/api"
	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	log.Println("DB Driver:", config.DBDriver)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	server, err := api.NewServer(config, db.NewStore(conn))
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
