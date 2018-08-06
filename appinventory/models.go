package appinventory

// Product ...
type Product struct {
	Sku         string
	ProductName string
	Quantity    int32
}

// InventoryIn ...
type InventoryIn struct {
	Sku              string
	ProductName      string
	ReceiptNumber    string
	Note             string
	OrderQuantity    int32
	ReceivedQuantity int32
	BuyPrice         int32
	Total            int32
}

// InventoryOut ...
type InventoryOut struct {
	Sku         string
	ProductName string
	Note        string
	OutQuantity int32
	SellPrice   int32
	Total       int32
}
