package appservice

import (
	rest "bitbucket.org/iandkenzt/inventories/restapi"
	"github.com/gorilla/mux"
)

// BuildAppServiceRoutes ...
func BuildAppServiceRoutes(r *mux.Router) {
	r.Methods("GET").Path("/is_alive").HandlerFunc(rest.AppSecretKeyMiddleware(IsAlive))
}
