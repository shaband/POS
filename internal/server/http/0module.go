package http

import (
	"go.uber.org/fx"

	"github.com/shaband/POS/internal/server/http/route"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), route.Module())
}
