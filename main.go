package main

import (
	"database/sql"
	"log"

	"github.com/aanhntm/restful-api/api"
	db "github.com/aanhntm/restful-api/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@172.17.0.2:5432/order?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}

	record := db.NewRecord(conn)
	server := api.NewServer(record)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
