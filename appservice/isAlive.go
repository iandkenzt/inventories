package appservice

import (
	"net/http"

	"github.com/iandkenzt/inventories/utils"
)

// Alive Struct
type Alive struct {
	IsAlive bool `json:"is_alive"`
}

// IsAlive Check machine status
func IsAlive(res http.ResponseWriter, req *http.Request) {

	var status Alive
	status.IsAlive = true

	utils.SendJSONResponse(res, 0, "Success", status, http.StatusOK)
}
