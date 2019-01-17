package appservice

import (
	"github.com/gorilla/mux"
)

// BuildAppServiceRoutes ...
func BuildAppServiceRoutes(r *mux.Router) {
	r.Methods("GET").Path("/is_alive").HandlerFunc(IsAlive)
}
