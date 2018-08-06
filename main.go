package main

import (
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"bitbucket.org/iandkenzt/inventories/restapi"
	"bitbucket.org/iandkenzt/inventories/utils"
	"github.com/gorilla/handlers"
)

func init() {
	// instantiate a logger
	utils.Logger.Out = os.Stdout
	utils.InitLogger()
}

func main() {

	port := "3000"

	if restapi.Conf.Port != "" {
		port = restapi.Conf.Port
	}

	utils.Logger.Info("Listening on port:", port)
	w := utils.Logger.Writer()

	apiRouter := RestAPIRouter()
	loggedRouter := handlers.LoggingHandler(w, apiRouter)

	http.ListenAndServe(":"+port, loggedRouter)

}
