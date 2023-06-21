package env

import (
	"os"
)

var (
	OracleUser     string
	OraclePassword string
	OracleHost     string
)

func Init() {

	os.Setenv("ORACLE_USER", "system")
	os.Setenv("ORACLE_PASSWORD", "tiger")
	os.Setenv("ORACLE_HOST", "localhost:1521/orcl")

	OracleUser = os.Getenv("ORACLE_USER")
	OraclePassword = os.Getenv("ORACLE_PASSWORD")
	OracleHost = os.Getenv("ORACLE_HOST")

}
