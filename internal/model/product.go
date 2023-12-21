package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"products"                                      swaggerignore:"true"`
	ID            int       `bun:",pk,autoincrement"                             example:"1"                         json:"id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"   swaggerignore:"true"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"   swaggerignore:"true"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"                         swaggerignore:"true"`
	Name          string    `example:"category1"                                 json:"name"`
	Code          string    `bun:",nullzero,notnull"                             example:"123"                       json:"code"`
	CostPrice     string    `bun:",nullzero,notnull"                             example:"10"                        json:"cost_price"`
	SellPrice     string    `bun:",nullzero,notnull"                             example:"100"                       json:"sell_price"`
	Image         string    `bun:",nullzero,notnull"                             example:"http://lorempicsm/100.jpg" json:"image"`
	CategoryID    int       `bun:",nullzero,notnull"                             example:"1"                         json:"category_id"`
	Category      *Category `bun:"rel:belongs-to,join:categories_category_id=id"`
}
