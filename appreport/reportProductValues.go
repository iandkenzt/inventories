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

	humanize "github.com/dustin/go-humanize"
	rest "github.com/iandkenzt/inventories/restapi"
	"github.com/iandkenzt/inventories/utils"
)

// CsvReportProductValues ...
func CsvReportProductValues(res http.ResponseWriter, req *http.Request) {

	// create connection
	database, err := sql.Open("sqlite3", rest.Conf.DbSqlite)
	if err != nil {
		println("Error db")
		log.Fatal(err)
	}
	defer database.Close()

	// create file csv
	fileName := fmt.Sprintf("./csv/REPORT_VALUE_PRODUCT_%s.csv", time.Now().Format("2006-01-02T150405"))
	path, _ := utils.CreateFile(fileName)

	// open file
	csvFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	defer csvFile.Close()

	// csv writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// summary of value produtcs
	rows, _ := database.Query(`
		WITH summary AS
			(
				SELECT sku, product_name, count(sku) as quantity,
					sum(total)/sum(received_quantity) as average_buy_price,
					count(sku)*(sum(total)/sum(received_quantity)) as total
				FROM inventory_in
				GROUP BY sku
			)
		SELECT count(sku) as total_sku, sum(quantity) as total_quantity, sum(total) as total_value
		FROM summary`)
	defer rows.Close()

	for rows.Next() {
		var totalSku, quantity, totalValue int
		rows.Scan(&totalSku, &quantity, &totalValue)

		// set header
		currentTime := time.Now()
		writer.Write([]string{"LAPORAN NILAI BARANG"})
		writer.Write([]string{"Tanggal Cetak", currentTime.Format("08 January 2006")})
		writer.Write([]string{"Jumlah SKU", strconv.Itoa(totalSku)})
		writer.Write([]string{"Jumlah Total Barang", strconv.Itoa(quantity)})
		writer.Write([]string{"Total Nilai", "Rp." + humanize.Comma(int64(totalValue))})
		writer.Write([]string{""})
	}
	defer rows.Close()

	data := []string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"}
	writer.Write(data)

	rows, _ = database.Query(`
		SELECT sku, product_name, count(sku) as quantity, 
		sum(total)/sum(received_quantity) as average_buy_price, 
		count(sku)*(sum(total)/sum(received_quantity)) as total
		FROM inventory_in
		GROUP BY sku`)
	defer rows.Close()

	for rows.Next() {
		var quantity, averageBuyPrice, total int
		var sku, productName string

		rows.Scan(&sku, &productName, &quantity, &averageBuyPrice, &total)

		data := []string{string(sku), string(productName), strconv.Itoa(quantity),
			"Rp." + humanize.Comma(int64(averageBuyPrice)), "Rp." + humanize.Comma(int64(total))}
		writer.Write(data)
	}
	defer rows.Close()
	writer.Flush()

	utils.SendJSONResponse(res, 0, "Success", fileName, http.StatusOK)

}
