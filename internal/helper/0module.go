package helper

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("helper", fx.Invoke())
}
