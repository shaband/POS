package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Inventory struct {
	bun.BaseModel `bun:"inventories"                                     swaggerignore:"true"`
	ID            int                `bun:",pk,autoincrement"                               example:"1"          json:"id"`
	CreatedAt     time.Time          `bun:",nullzero,notnull,default:current_timestamp"     swaggerignore:"true"`
	UpdatedAt     time.Time          `bun:",nullzero,notnull,default:current_timestamp"     swaggerignore:"true"`
	DeletedAt     time.Time          `bun:",soft_delete,nullzero"                           swaggerignore:"true"`
	Name          string             `example:"category1"                                   json:"name"`
	Address       string             `bun:",nullzero,notnull"                               example:"address1"   json:"address"`
	Products      []Product          `bun:"m2m:inventory_to_product,join:Inventory=Product"`
	ProductStock  InventoryToProduct `bun:"rel:has-one,join:id=inventory_id"`
}
