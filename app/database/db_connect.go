package database

import (
	"database/sql"
	_ "github.com/golang-migrate/migrate/source/file"
	go_ora "github.com/sijms/go-ora/v2"
	"golang-blueprint-clean/app/env"
	"strconv"
)

func ConnectDB() *sql.DB {

	port, err := strconv.Atoi(env.OraclePort)
	if err != nil {
		panic(err.Error())
	}

	databaseURL := go_ora.BuildUrl(env.OracleHost, port, env.OracleService, env.OracleUser, env.OraclePassword, nil)

	db, err := sql.Open("oracle", databaseURL)
	if err != nil {

		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		db.Close()

	}

	return db
}
