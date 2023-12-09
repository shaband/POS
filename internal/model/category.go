package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"categories"                                  swaggerignore:"true"`
	ID            int       `bun:",pk,autoincrement"                           example:"1"          json:"id"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"                       swaggerignore:"true"`
	Name          string    `example:"category1"                               json:"name"`
	CategoryID    int       `example:"1"                                       json:"category_id"`
}
