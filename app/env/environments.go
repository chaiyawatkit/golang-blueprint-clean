package env

import (
	"os"
)

var (
	AppFeatureServiceUrl string
)

func Init() {

	os.Setenv("AppFeatureServiceUrl", "http://localhost:8080")

	AppFeatureServiceUrl = os.Getenv("AppFeatureServiceUrl")

}
