package infra

import (
	"go.uber.org/fx"

	"github.com/shaband/POS/internal/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
