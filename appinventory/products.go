package appinventory

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"bitbucket.org/iandkenzt/inventories/utils"
)

// StoreStockProduct ...
func StoreStockProduct(res http.ResponseWriter, req *http.Request) {

	var err error
	var temp interface{}
	var p Product

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&temp)

	if err != nil {
		utils.SendJSONResponse(res, 1, "Error", nil, http.StatusBadRequest)
		return
	}

	jsonData := temp.(map[string]interface{})

	// checking data
	if _, ok := jsonData["sku"]; !ok {
		utils.SendJSONResponse(res, 1, "Error", "data is not valid, error found!", http.StatusBadRequest)
		return
	}

	if _, ok := jsonData["product_name"]; !ok {
		utils.SendJSONResponse(res, 1, "Error", "data is not valid, error found!", http.StatusBadRequest)
		return
	}

	// set product
	p.Sku = jsonData["sku"].(string)
	p.ProductName = jsonData["product_name"].(string)
	p.Quantity = 0

	if _, ok := jsonData["quantity"]; ok {
		p.Quantity = int32(jsonData["quantity"].(float64))
	}

	// open connection to sqlite
	database, err := sql.Open("sqlite3", string(Conf.DbSqlite))
	if err != nil {
		println("Error open connection to sqlite")
	}
	defer database.Close()

	// store data product
	stmt, _ := database.Prepare("INSERT INTO products(sku, product_name, quantity) VALUES(?, ?, ?)")
	stmt.Exec(p.Sku, p.ProductName, p.Quantity)
	defer stmt.Close()

	utils.SendJSONResponse(res, 0, "Success", p, http.StatusOK)

}

// GetReportProducts ...
func GetReportProducts(res http.ResponseWriter, req *http.Request) {

	var products []map[string]interface{}

	// connection to sqlite
	db, err := sql.Open("sqlite3", Conf.DbSqlite)
	if err != nil {
		println("Error connect to DB")
		utils.SendJSONResponse(res, 1, "Error", "Error connect to DB", http.StatusBadRequest)
		return
	}
	defer db.Close()

	// create response data products
	rows, err := db.Query("SELECT sku,product_name,quantity FROM products")
	if err != nil {
		println("Error query data to DB")
		utils.SendJSONResponse(res, 1, "Error", "Error query data to DB", http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var quantity int
		var sku, productName string

		rows.Scan(&sku, &productName, &quantity)

		data := map[string]interface{}{"sku": sku, "product_name": productName, "quantity": quantity}
		products = append(products, data)
	}
	defer rows.Close()

	utils.SendJSONResponse(res, 0, "Success", products, http.StatusOK)

}
