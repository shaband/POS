package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"posts" swaggerignore:"true"`
	ID            int           `bun:",pk,autoincrement" json:"id" example:"1"`
	CreatedAt     time.Time     `bun:",nullzero,notnull,default:current_timestamp" `
	UpdatedAt     time.Time     `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	DeletedAt     time.Time     `bun:",soft_delete,nullzero" swaggerignore:"true"`
	Name          string        `json:"name"  example:"category1"`
	CategoryID    int           `json:"category_id"  example:"1"`
	Parent        interface{}   `bun:"rel:belongs-to,join:category_id=id" swaggerignore:"true"`
	Children      []interface{} `bun:"rel:has-many,join:id=category_id" swaggerignore:"true"`
}
