package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/florian-nguyen/golang-training/tech-school/simple-bank/db/util"
	_ "github.com/lib/pq" // blank identifier required to avoid automatic delete on save
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
