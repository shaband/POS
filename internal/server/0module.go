package server

import (
	"go.uber.org/fx"

	"github.com/shaband/POS/internal/server/http"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
