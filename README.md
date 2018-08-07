# inventories

RestAPI service for inventory activity

## How to run With Native Go

* Run Package
> `go run ./*.go`

## Endpoint ##

### Check service Rest API
```
curl --request GET \
  --url http://localhost:3000/api/v1/is_alive \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":{"is_alive":true}}
```

### Check service Rest API
```
curl --request GET \
  --url http://localhost:3000/api/v1/is_alive \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":{"is_alive":true}}
```

### Get data product
```
curl --request GET \
  --url http://localhost:3000/api/v1/product \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":[{"product_name":"Zalekia Plain Casual Blouse (L,Broken White)","quantity":154,"sku":"SSI-D00791015-LL-BWH"},{"product_name":"Zalekia Plain Casual Blouse (M,Broken White)","quantity":138,"sku":"SSI-D00791077-MM-BWH"}]}
```

### Insert data product
```
curl --request PUT \
  --url http://localhost:3000/api/v1/product/insert \
  --data '{"sku": "TEST-PRODUCT-002","product_name": "TEST-PRODUCT-002","quantity": 100}' \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":{"Sku":"TEST-PRODUCT-002","ProductName":"TEST-PRODUCT-002","Quantity":100}}
```

### Insert data Inventory In
```
curl --request POST \
  --url http://localhost:3000/api/v1/inventory/in \
  --data '{"sku": "TEST-INVENTORY-IN-001","product_name": "Product tester TEST-INVENTORY-IN-001","receipt_number": "TEST-INVENTORY-IN-001","note": "TEST-INVENTORY-IN-001","order_quantity": 1,"received_quantity":1,"buy_price": 100000}' \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":[""]}
```

### Insert data Inventory Out
```
curl --request POST \
  --url http://localhost:3000/api/v1/inventory/out \
  --data '{"sku": "TEST-PRODUCT-INVENTORY-OUT-001","product_name": "TEST-PRODUCT-INVENTORY-OUT-001","note": "TEST-PRODUCT-INVENTORY-OUT-001","time": 10000,"sell_price": 10000,"out_quantity": 100}' \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":[]}
```

### Generate CSV Report Product
```
curl --request GET \
  --url http://localhost:3000/api/v1/report/products \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":"./csv/REPORT_PRODUCT_2018-08-07T165910.csv"}
```

### Generate CSV Report Product values
```
curl --request GET \
  --url http://localhost:3000/api/v1/report/product_values \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":"./csv/REPORT_VALUE_PRODUCT_2018-08-07T170013.csv"}
```

### Generate CSV Report Selling Product
```
curl --request GET \
  --url http://localhost:3000/api/v1/report/selling_products \
  -H 'X-Api-Key: 99saf8le-s7td-408c-K2iN-a3fv69encTry'
```
```
Response: {"error":0,"message":"Success","data":"./csv/REPORT_SELLING_PRODUCT_2018-08-07T170052.csv"}
```