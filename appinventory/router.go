package appinventory

import (
	rest "bitbucket.org/iandkenzt/inventories/restapi"

	"github.com/gorilla/mux"
)

// BuildAppInventoryRoutes ...
func BuildAppInventoryRoutes(r *mux.Router) {
	// products
	r.Methods("GET").Path("/product").HandlerFunc(rest.AppSecretKeyMiddleware(GetReportProducts))
	r.Methods("PUT").Path("/product/insert").HandlerFunc(rest.AppSecretKeyMiddleware(StoreStockProduct))
}
