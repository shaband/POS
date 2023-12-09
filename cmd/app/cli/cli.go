package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/shaband/POS/internal/app"
	"github.com/shaband/POS/internal/app/appcontext"
)

func Start(module fx.Option) {
	_ = app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
