package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Client struct {
	bun.BaseModel `bun:"clients" swaggerignore:"true"`

	ID         int       `bun:",pk,autoincrement"                           json:"id"`
	CreatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt  time.Time `bun:",soft_delete,nullzero"`
	Name       string    `bun:",notnull,nullzero"                           json:"name"`
	Email      string    `bun:",notnull,nullzero"                           json:"email"`
	Phone      string    `bun:",notnull,nullzero"                           json:"phone"`
	IsSupplier bool      `bun:",notnull,nullzero,default:true"              json:"is_supplier"`
	IsClient   bool      `bun:",notnull,nullzero,default:false"             json:"is_client"`
	// Image     string    `json:"image",nullzero`
}
