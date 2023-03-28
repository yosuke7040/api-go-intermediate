package repositories_test

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}
