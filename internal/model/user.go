package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users" swaggerignore:"true"`

	ID        int       `bun:",pk,autoincrement"                           json:"id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
	Name      string    `bun:",notnull"                                    json:"name"`
	Email     string    `bun:",notnull"                                    json:"email"`
	Phone     string    `bun:",notnull"                                    json:"phone"`
	Password  string    `bun:",notnull"                                    json:"password"`
	Image     string    `json:"image"`
}
