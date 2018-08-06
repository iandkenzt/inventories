package restapi

import (
	"os"

	"github.com/joho/godotenv"
)

// Configuration Env Struct
type Configuration struct {
	Port         string
	AppSecretKey string
	APIPrefix    string
	APIVersion   string
	DbSqlite     string
}

// Conf Var
var Conf Configuration

// LoadEnv ...
func init() {
	if Conf == (Configuration{}) {
		godotenv.Load()

		// APP Config
		Conf.Port = os.Getenv("PORT")
		Conf.AppSecretKey = os.Getenv("APP_SECRET_KEY")
		Conf.APIPrefix = os.Getenv("API_PREFIX")
		Conf.APIVersion = os.Getenv("API_VERSION")
		Conf.DbSqlite = os.Getenv("DB_SQLITE")
	}
}
