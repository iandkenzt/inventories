package appreport

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	rest "bitbucket.org/iandkenzt/inventories/restapi"
	"bitbucket.org/iandkenzt/inventories/utils"
)

// CsvReportSellingProduct ...
func CsvReportSellingProduct(res http.ResponseWriter, req *http.Request) {

	database, err := sql.Open("sqlite3", rest.Conf.DbSqlite)
	if err != nil {
		println("Error db")
		log.Fatal(err)
	}
	defer database.Close()

	// create file csv
	fileName := fmt.Sprintf("./csv/REPORT_SELLING_PRODUCT_%s.csv", time.Now().Format("2006-01-02T150405"))
	path, _ := utils.CreateFile(fileName)

	csvFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// set header
	data := []string{"ID Pesanan", "Waktu", "SKU", "Nama Barang",
		"Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"}
	writer.Write(data)

	rows, _ := database.Query(`
		WITH summary AS
		(
			SELECT product_name, sum(total)/sum(received_quantity) as buy_price
			FROM inventory_in
			GROUP BY product_name
		)
		SELECT sell_id, time, sku, summary.product_name, out_quantity, sell_price, 
			total, buy_price, total-(buy_price*out_quantity) as profit
		FROM inventory_out
		LEFT JOIN summary ON inventory_out.product_name = summary.product_name	
	`)
	defer rows.Close()

	for rows.Next() {
		var sellID, time, total, outQuantity, sellPrice, buyPrice, profit int
		var sku, productName string

		rows.Scan(&sellID, &time, &sku, &productName, &outQuantity, &sellPrice, &total, &buyPrice, &profit)

		println(sellID, time, total, outQuantity, sellPrice, buyPrice, profit, sku, productName)

		data := []string{
			strconv.Itoa(sellID), strconv.Itoa(time), string(sku), string(productName),
			strconv.Itoa(outQuantity), strconv.Itoa(sellPrice), strconv.Itoa(total),
			strconv.Itoa(buyPrice), strconv.Itoa(profit)}
		writer.Write(data)
	}
	defer rows.Close()
	writer.Flush()

	utils.SendJSONResponse(res, 0, "Success", fileName, http.StatusOK)

}
