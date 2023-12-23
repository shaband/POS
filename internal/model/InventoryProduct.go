package model

import "github.com/uptrace/bun"

type InventoryToProduct struct {
	bun.BaseModel     `bun:"table:inventory_to_product" swaggerignore:"true"`
	ProductID         int        `bun:",pk"`
	Product           *Product   `bun:"rel:belongs-to,join:product_id=id"`
	InventoryID       int        `bun:",pk"`
	Inventory         *Inventory `bun:"rel:belongs-to,join:inventory_id=id"`
	Amount            int        `json:"amount"`
	SellInvoicesCount int        `json:"sell_invoices_count"`
	BuyInvoicesCount  int        `json:"buy_invoices_count"`
}
