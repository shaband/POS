package dto

type InventoryDTO struct {
	Name    string `example:"inventory1" json:"name"    valid:"type(string),required~name field is required"`
	Address string `example:"address1"   json:"address" valid:"type(string),required~address field is required"`
}

type InventoryToProductDTO struct {
	ProductID         int `json:"product_id"`
	InventoryID       int `json:"inventory_id"`
	Amount            int `json:"amount"`
	SellInvoicesCount int `json:"sell_invoices_count"`
	BuyInvoicesCount  int `json:"buy_invoices_count"`
}
