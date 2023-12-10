package helper

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("helper", fx.Provide(
		validator.New,
		// RegisterValidator,
	),
	)
}
