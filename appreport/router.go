package appreport

import (
	rest "github.com/iandkenzt/inventories/restapi"

	"github.com/gorilla/mux"
)

// BuildAppReportsRoutes ...
func BuildAppReportsRoutes(r *mux.Router) {

	r.Methods("GET").Path("/report/products").HandlerFunc(rest.AppSecretKeyMiddleware(CsvReportProducts))
	r.Methods("GET").Path("/report/product_values").HandlerFunc(rest.AppSecretKeyMiddleware(CsvReportProductValues))
	r.Methods("GET").Path("/report/selling_products").HandlerFunc(rest.AppSecretKeyMiddleware(CsvReportSellingProduct))

}
