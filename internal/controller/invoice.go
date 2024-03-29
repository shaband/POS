package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/shaband/POS/internal/helper"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/output"
	"github.com/shaband/POS/internal/service"
)

type Invoice struct {
	fx.In
	Validator *helper.Validator
	Service   *service.Invoice
	Route     fiber.Router `name:"api-v1"`
}

func RegisterInvoice(c Invoice) {
	c.Route.Get("/invoices", c.GetInvoices)
	c.Route.Post("/invoices", c.AddInvoice)
}

// GetInvoices godoc
//
//	@Summary		get All Invoices
//	@Description	get All Invoices
//	@Tags			invoices
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	output.Invoice
//	@Router			/invoices [get]
func (c *Invoice) GetInvoices(ctx *fiber.Ctx) error {
	invoices, err := c.Service.GetInvoices(ctx.Context())
	if err != nil {
		return err
	}
	result := []*output.Invoice{}

	for _, invoice := range invoices {
		result = append(result, &output.Invoice{
			ID: uint(invoice.ID),
			Client: output.NameWithID{
				ID:   invoice.ClientID,
				Name: invoice.Client.Name,
			},
			Inventory: output.NameWithID{
				ID:   invoice.InventoryID,
				Name: invoice.Inventory.Name,
			},
			User: output.NameWithID{
				ID:   invoice.UserID,
				Name: invoice.User.Name,
			},
		})
	}

	return ctx.JSON(result)
}

// AddCategories godoc
//
//	@Summary		Add New Invoice
//	@Description	Add New Invoice
//	@Tags			invoices
//	@Accept			json
//	@Produce		json
//	@param			request	body		dto.InvoiceDTO	true	"invoice data"
//	@Success		200		{string}	string			"Invoice Stored Successfully"
//
//	@Failure		400		{object}	output.HTTPError
//	@Router			/invoices [Post]
func (c *Invoice) AddInvoice(ctx *fiber.Ctx) error {
	InvoiceDTO := new(dto.InvoiceDTO)
	_ = ctx.BodyParser(InvoiceDTO)

	isValid, _ := c.Validator.Validate(InvoiceDTO)
	if isValid {
		// TODO:: check if not sell  cost and price is required
		if !InvoiceDTO.IsSell {
			m := make(map[string]interface{})
			ctx.BodyParser(&m)
			// items:=m["items"]
			for i := 0; i < len(InvoiceDTO.Items); i++ {
				// c.Validator.ValidateMap(items[i], map[string]interface{}{
				// 	"unit_sell_price":"required,numeric",
				// 	"unit_cost_price":"required,numeric",
				// 	"amount":"required,numeric",
				// 	"product_id":"required",

				// }	)
			}
			// 		c.Validator.ValidateMap(map[string]interface{}{
			// 			"items": []map[string]interface{}{
			// 				"line1":       "",
			// 				"line2":       "",
			// 				"postal-code": "",
			// 			},
			// 		}, map[string]interface{}{
			// 			"items": []map[string]interface{}{
			// 				"line1":       "required,alphanum",
			// 				"line2":       "alphanum",
			// 				"postal-code": "numeric",
			// 			},
			// 		})
		}
	}
	_, err := c.Service.AddInvoice(ctx.Context(), InvoiceDTO)
	if err == nil {
		ctx.WriteString("Invoice Stored Successfully")
	}
	// }
	return err
}
