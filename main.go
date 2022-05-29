package main

import (
	"database/sql"
	"log"

	"github.com/aanhntm/restful-api/api"
	db "github.com/aanhntm/restful-api/db/sqlc"
	"github.com/aanhntm/restful-api/util"
	_ "github.com/lib/pq"
)

/*const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/order?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)*/

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}

	record := db.NewRecord(conn)
	server := api.NewServer(record)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

	//fmt.Println(config)
}
