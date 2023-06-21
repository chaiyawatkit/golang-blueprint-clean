package database

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	_ "github.com/golang-migrate/migrate/source/file"
	"golang-blueprint-clean/app/env"
)

func ConnectDB() *sql.DB {

	connection := fmt.Sprintf(
		"user=%s password=%s connectString=%s",
		env.OracleUser,
		env.OraclePassword,
		env.OracleHost,
	)

	db, err := sql.Open("godror", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
