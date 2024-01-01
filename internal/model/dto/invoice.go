package dto

type InvoiceItemDTO struct {
	UnitSellPrice float64 `example:"10000" json:"unit_sell_price" valid:"required"`
	UnitCostPrice float64 `example:"1000"  json:"unit_cost_price" valid:"required"`
	Amount        float64 `example:"1"     json:"amount"          valid:"required"`
	ProductID     int     `example:"1"     json:"product_id"      valid:"required"`
}

type InvoiceDTO struct {
	IsSell      bool    `example:"true"  json:"is_sell"      valid:"type(bool)"`
	TotalCost   float64 `example:"1000"  json:"total_cost"   valid:"-"`
	TotalPrice  float64 `example:"10000" json:"total_price"  valid:"-"`
	InventoryID int     `example:"6"     json:"inventory_id" valid:"required"`
	ClientID    int     `example:"1"     json:"client_id"    valid:"required"`
	UserID      int     `example:"1"     json:"user_id"      valid:"required"`
	Items       []InvoiceItemDTO
}
