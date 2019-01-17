package appinventory

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iandkenzt/inventories/utils"
)

// StoreInventoryOut ...
func StoreInventoryOut(res http.ResponseWriter, req *http.Request) {

	var err error
	var temp interface{}
	var p InventoryOut

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&temp)

	if err != nil {
		utils.SendJSONResponse(res, 1, "Error", nil, http.StatusBadRequest)
		return
	}

	jsonData := temp.(map[string]interface{})

	// validation data
	if _, ok := jsonData["sku"]; !ok {
		utils.SendJSONResponse(res, 1, "Data is not valid, sku not found or null value!", []string{}, http.StatusBadRequest)
		return
	}

	if _, ok := jsonData["product_name"]; !ok {
		utils.SendJSONResponse(res, 1, "Data is not valid, product_name not found or null value!", []string{}, http.StatusBadRequest)
		return
	}

	// set product
	p.Sku = jsonData["sku"].(string)
	p.ProductName = jsonData["product_name"].(string)
	p.Note = jsonData["note"].(string)
	p.OutQuantity = int32(jsonData["out_quantity"].(float64))
	p.SellPrice = int32(jsonData["sell_price"].(float64))
	p.Total = p.OutQuantity * p.SellPrice
	currentTime := time.Now()

	// open connection to sqlite
	database, err := sql.Open("sqlite3", string(Conf.DbSqlite))
	if err != nil {
		println("Error db")
	}
	defer database.Close()

	// store data inventory out
	stmt, _ := database.Prepare("INSERT INTO inventory_out(sku, time, product_name, out_quantity, sell_price, total, note) values(?, ?, ?, ?, ?, ?, ?)")
	stmt.Exec(p.Sku, currentTime.Format("2006-01-02 15:04:05"), p.ProductName, p.OutQuantity, p.SellPrice, p.Total, p.Note)
	defer stmt.Close()

	utils.SendJSONResponse(res, 0, "Success", []string{}, http.StatusOK)

}
