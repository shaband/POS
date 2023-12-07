package app

import (
	"go.uber.org/fx"

	"github.com/shaband/POS/internal/app/appconfig"
	"github.com/shaband/POS/internal/app/appcontext"
	"github.com/shaband/POS/internal/controller"
	"github.com/shaband/POS/internal/infra"
	"github.com/shaband/POS/internal/repo"
	"github.com/shaband/POS/internal/server"
	"github.com/shaband/POS/internal/service"
	"github.com/shaband/POS/internal/x/logger"
	"github.com/shaband/POS/internal/x/logger/fxlogger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		controller.Module(),
		infra.Module(),
		repo.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
