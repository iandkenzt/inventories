package appinventory

import (
	rest "github.com/iandkenzt/inventories/restapi"

	"github.com/gorilla/mux"
)

// BuildAppInventoryRoutes ...
func BuildAppInventoryRoutes(r *mux.Router) {
	// products
	r.Methods("GET").Path("/product").HandlerFunc(rest.AppSecretKeyMiddleware(GetReportProducts))
	r.Methods("PUT").Path("/product/insert").HandlerFunc(rest.AppSecretKeyMiddleware(StoreStockProduct))

	// inventory
	r.Methods("POST").Path("/inventory/in").HandlerFunc(rest.AppSecretKeyMiddleware(StoreInventoryIn))
	r.Methods("POST").Path("/inventory/out").HandlerFunc(rest.AppSecretKeyMiddleware(StoreInventoryOut))

}
