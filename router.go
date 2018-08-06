package main

import (
	"bitbucket.org/iandkenzt/inventories/appinventory"
	"bitbucket.org/iandkenzt/inventories/appservice"
	"bitbucket.org/iandkenzt/inventories/restapi"
	"github.com/gorilla/mux"
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

	return router

}
