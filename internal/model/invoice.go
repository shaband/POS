package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Invoice struct {
	ID          int            `bun:",pk,autoincrement"                           example:"1"          json:"id"`
	CreatedAt   time.Time      `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	UpdatedAt   time.Time      `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	DeletedAt   time.Time      `bun:",soft_delete,nullzero"                       swaggerignore:"true"`
	IsSell      bool           `json:"is_sell"`
	TotalCost   float64        `json:"total_cost"`
	TotalPrice  float64        `json:"total_price"`
	InventoryID int            `json:"inventory_id"`
	Inventory   Inventory      `bun:"rel:belongs-to,join:inventory_id=id"`
	ClientID    int            `json:"client_id"`
	Client      Client         `bun:"rel:belongs-to,join:client_id=id"`
	UserID      int            `json:"user_id"`
	User        User           `bun:"rel:belongs-to,join:user_id=id"`
	Items       []*InvoiceItem `bun:"rel:has-many,join:id=invoice_id"`
}
type InvoiceItem struct {
	bun.BaseModel `bun:"inventories"                                 swaggerignore:"true"`
	ID            int       `bun:",pk,autoincrement"                           example:"1"          json:"id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"                       swaggerignore:"true"`
	UnitSellPrice float64   `json:"unit_sell_price"`
	UnitCostPrice float64   `json:"unit_cost_price"`
	Amount        float64   `json:"amount"`
	TotalCost     float64   `json:"total_cost"`
	TotalPrice    float64   `json:"total_price"`
	// UserID        int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Product   Product   `bun:"rel:belongs-to,join:product_id=id"`
	InvoiceID int       `json:"invoice_id"`
	Invoice   Inventory `bun:"rel:belongs-to,join:invoice_id=id"`
}
