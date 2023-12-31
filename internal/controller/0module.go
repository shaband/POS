package controller

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("controller", fx.Invoke(
		RegisterUser,
		RegisterSwagger,
		RegisterCategory,
		RegisterClient,
		RegisterProduct,
		RegisterInventory,
		RegisterInvoice,
	))
}
