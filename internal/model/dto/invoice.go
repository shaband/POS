package dto

type InvoiceItemDTO struct {
	UnitSellPrice float64 `json:"unit_sell_price" swaggerignore:"true" valid:"required"`
	UnitCostPrice float64 `json:"unit_cost_price" swaggerignore:"true" valid:"required"`
	Amount        float64 `json:"amount"          valid:"required"`
	ProductID     int     `json:"product_id"      valid:"required"`
}

type InvoiceDTO struct {
	IsSell      bool    `json:"is_sell"      valid:"required"`
	TotalCost   float64 `json:"total_cost"   valid:"-"`
	TotalPrice  float64 `json:"total_price"  valid:"-"`
	InventoryID int     `json:"inventory_id" valid:"required"`
	ClientID    int     `json:"client_id"    valid:"required"`
	UserID      int     `json:"user_id"      swaggerignore:"true" valid:"required"`
	Items       []InvoiceItemDTO
}
