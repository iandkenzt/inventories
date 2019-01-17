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

	rest "github.com/iandkenzt/inventories/restapi"
	"github.com/iandkenzt/inventories/utils"
)

// CsvReportProducts ...
func CsvReportProducts(res http.ResponseWriter, req *http.Request) {

	// connection to sqlite
	database, err := sql.Open("sqlite3", rest.Conf.DbSqlite)
	if err != nil {
		println("Error db")
		log.Fatal(err)
	}
	defer database.Close()

	// create file csv
	fileName := fmt.Sprintf("./csv/REPORT_PRODUCT_%s.csv", time.Now().Format("2006-01-02T150405"))
	path, _ := utils.CreateFile(fileName)

	// open file
	csvFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	defer csvFile.Close()

	// csv write
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{""})

	// set header
	rows, _ := database.Query("SELECT count(sku) AS total_product FROM products")
	for rows.Next() {
		var totalProducts int
		rows.Scan(&totalProducts)

		currentTime := time.Now()
		writer.Write([]string{"LAPORAN PRODUK"})
		writer.Write([]string{"Tanggal Cetak", currentTime.Format("08 January 2006")})
		writer.Write([]string{"Jumlah Produk Item", strconv.Itoa(totalProducts)})
		writer.Write([]string{""})
	}
	defer rows.Close()

	data := []string{"SKU", "Nama Item", "Jumlah Sekarang"}
	writer.Write(data)

	// write data products
	rows, _ = database.Query("SELECT sku,product_name,quantity FROM products")
	for rows.Next() {
		var quantity int
		var sku, productName string

		rows.Scan(&sku, &productName, &quantity)

		data := []string{string(sku), string(productName), strconv.Itoa(quantity)}
		writer.Write(data)
	}
	defer rows.Close()
	writer.Flush()

	utils.SendJSONResponse(res, 0, "Success", path, http.StatusOK)

}
