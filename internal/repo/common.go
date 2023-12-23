package repo

import "github.com/uptrace/bun"

func selectNameAndID(sq *bun.SelectQuery) *bun.SelectQuery {
	return sq.Column("id").Column("name")
}
