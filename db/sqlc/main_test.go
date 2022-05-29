package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/aanhntm/restful-api/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
