package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaband/POS/internal/helper"
	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/output"
	"github.com/shaband/POS/internal/service"
	"go.uber.org/fx"

)

type Inventory struct {
	fx.In
	Validator *helper.Validator
	Service   *service.Inventory
	Route     fiber.Router `name:"api-v1"`
}

func RegisterInventory(c Inventory) {
	c.Route.Get("/inventories", c.GetInventories)
	c.Route.Post("/inventories", c.AddInventory)
	c.Route.Patch("/inventories/:id", c.UpdateInventory)
	c.Route.Delete("/inventories/:id", c.DeleteInventory)
}

// GetInventories godoc
//
//	@Summary		get All Inventories
//	@Description	get All Inventories
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	output.Inventory
//	@Router			/inventories [get]
func (c Inventory) GetInventories(ctx *fiber.Ctx) error {
	inventories, err := c.Service.GetInventories(ctx.Context())
	if err == nil {
		var results []output.Inventory = []output.Inventory{}
		for _, inventory := range inventories {
			results = append(results, output.Inventory{
				ID:      uint(inventory.ID),
				Name:    inventory.Name,
				Address: inventory.Address,
			})
		}
		return ctx.JSON(results)
	}
	return err
}

// AddInventories godoc
//
//	@Summary		Add New Inventory
//	@Description	Add New Inventory
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//	@param			request	body		dto.InventoryDTO	true	"inventory data"
//	@Success		200		{object}	output.Inventory
//
//	@Failure		400		{object}	output.HTTPError
//	@Router			/inventories [Post]
func (c Inventory) AddInventory(ctx *fiber.Ctx) error {
	InventoryDTO := new(dto.InventoryDTO)
	_ = ctx.BodyParser(InventoryDTO)
	isValid, err := c.Validator.Validate(InventoryDTO)
	if isValid {
		results, err := c.Service.AddInventory(ctx.Context(), InventoryDTO)
		if err == nil {
			return ctx.JSON(output.Inventory{
				ID:      uint(results.ID),
				Name:    results.Name,
				Address: results.Address,
			})
		}
	}
	return err
}

// UpdateInventories godoc
//
//	@Summary		Update  Inventory
//	@Description	Update  Inventory
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//	@param			request			body		dto.InventoryDTO	true	"inventory data"
//	@param			inventory_id	path		int					true	"inventory id"
//
//	@Success		200				{object}	model.Inventory
//
//	@Failure		400				{object}	output.HTTPError
//	@Router			/inventories/{inventory_id} [Patch]
func (c Inventory) UpdateInventory(ctx *fiber.Ctx) error {
	InventoryDTO := new(dto.InventoryDTO)
	_ = ctx.BodyParser(InventoryDTO)

	isValid, err := c.Validator.Validate(InventoryDTO)
	if isValid {
		id, err := ctx.ParamsInt("id")
		if err == nil {
			var results *model.Inventory
			results, err = c.Service.UpdateInventory(ctx.Context(), id, InventoryDTO)
			if err != nil {
				return err
			}
			return ctx.JSON(output.Inventory{
				ID:      uint(results.ID),
				Name:    results.Name,
				Address: results.Address,
			})
		}
	}
	return err
}

// DELETEInventories godoc
//
//	@Summary		delete  Inventory
//	@Description	delete  Inventory
//	@Tags			inventories
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int		true	"Inventory ID"
//	@Success		200	{string}	string	"deleted successfully"
//
//	@Router			/inventories/{id} [Delete]
func (c Inventory) DeleteInventory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err == nil {
		err = c.Service.DeleteInventory(ctx.Context(), id)
		if err == nil {
			_, err = ctx.WriteString("deleted successfully")
		}
	}

	return err
}
