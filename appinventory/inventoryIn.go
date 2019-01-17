package appinventory

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iandkenzt/inventories/utils"
)

// StoreInventoryIn ...
func StoreInventoryIn(res http.ResponseWriter, req *http.Request) {

	var err error
	var temp interface{}
	var invIn InventoryIn

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

	// open connection to sqlite
	database, err := sql.Open("sqlite3", string(Conf.DbSqlite))
	if err != nil {
		println("Error db")
	}
	defer database.Close()

	// set value struct
	invIn.ProductName = jsonData["product_name"].(string)

	// check sku product exsits or not
	stmt, err := database.Prepare("SELECT * FROM products WHERE product_name = (?)")
	stmt.Exec(invIn.ProductName)
	defer stmt.Close()

	var productName string
	var productQuantity int32
	err = stmt.QueryRow(invIn.ProductName).Scan(&productName, &productQuantity)
	if err == nil {
		utils.SendJSONResponse(res, 0, "product_name not found or null value!", []string{}, http.StatusBadRequest)
		return
	}

	// set value struct
	invIn.Sku = jsonData["sku"].(string)
	invIn.ReceiptNumber = jsonData["receipt_number"].(string)
	invIn.Note = jsonData["note"].(string)
	invIn.OrderQuantity = int32(jsonData["order_quantity"].(float64))
	invIn.ReceivedQuantity = int32(jsonData["received_quantity"].(float64))
	invIn.BuyPrice = int32(jsonData["buy_price"].(float64))
	invIn.Total = int32(invIn.OrderQuantity * invIn.BuyPrice)
	currentTime := time.Now()

	// store data inventory in
	queryBuilder := `
	INSERT INTO inventory_in(
		sku, product_name, receipt_number, note, time, order_quantity, 
		received_quantity, buy_price, total)
	values(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, _ = database.Prepare(queryBuilder)
	stmt.Exec(invIn.Sku, invIn.ProductName, invIn.ReceiptNumber, invIn.Note, currentTime.Format("2006-01-02 15:04:05"), invIn.OrderQuantity, invIn.ReceivedQuantity, invIn.BuyPrice, invIn.Total)
	defer stmt.Close()

	productQuantity = productQuantity + invIn.OrderQuantity
	stmt, _ = database.Prepare("UPDATE products SET quantity=(?) WHERE product_name=(?)")
	stmt.Exec(productQuantity, productName)
	defer stmt.Close()

	utils.SendJSONResponse(res, 0, "Success", []string{""}, http.StatusOK)

}
