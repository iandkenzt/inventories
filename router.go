package main

import (
	"github.com/gorilla/mux"
	"github.com/iandkenzt/inventories/appinventory"
	"github.com/iandkenzt/inventories/appreport"
	"github.com/iandkenzt/inventories/appservice"
	"github.com/iandkenzt/inventories/restapi"
)

// RestAPIRouter ...
func RestAPIRouter() *mux.Router {

	apiPrefix := restapi.Conf.APIPrefix
	apiVersion := restapi.Conf.APIVersion

	router := mux.NewRouter()
	restAPIRouter := router.PathPrefix(apiPrefix + apiVersion).Subrouter()

	// registration blueprint route of API Apps
	appservice.BuildAppServiceRoutes(restAPIRouter)
	appinventory.BuildAppInventoryRoutes(restAPIRouter)
	appreport.BuildAppReportsRoutes(restAPIRouter)

	return router

}
