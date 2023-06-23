package env

import (
	"os"
)

var (
	OracleUser     string
	OraclePassword string
	OracleHost     string
	OraclePort     string
	OracleService  string
)

func Init() {

	os.Setenv("ORACLE_USER", "system")
	os.Setenv("ORACLE_PASSWORD", "tiger")
	os.Setenv("ORACLE_HOST", "localhost")
	os.Setenv("ORACLE_PORT", "1521")
	os.Setenv("ORACLE_SERVICE", "orcl")

	OracleUser = os.Getenv("ORACLE_USER")
	OraclePassword = os.Getenv("ORACLE_PASSWORD")
	OracleHost = os.Getenv("ORACLE_HOST")
	OraclePort = os.Getenv("ORACLE_PORT")
	OracleService = os.Getenv("ORACLE_SERVICE")

}
