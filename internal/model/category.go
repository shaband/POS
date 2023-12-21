package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"categories"                                  swaggerignore:"true"`
	ID            int        `bun:",pk,autoincrement"                           example:"1"          json:"id"`
	CreatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	UpdatedAt     time.Time  `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	DeletedAt     time.Time  `bun:",soft_delete,nullzero"                       swaggerignore:"true"`
	Name          string     `bun:",nullzero,notnull"                           example:"category1"  json:"name"`
	CategoryID    int        `bun:",nullzero,notnull"                           example:"1"          json:"category_id"`
	Products      []*Product `bun:"rel:has-many,join:id=category_id"`
}
