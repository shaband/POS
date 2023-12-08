package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/service"
	"go.uber.org/fx"
)

type Category struct {
	fx.In
	Service *service.Category
	Route   fiber.Router `name:"api-v1"`
}

func RegisterCategory(c Category) {
	c.Route.Get("/categories", c.GetCategories)
	c.Route.Post("/categories", c.AddCategory)
	c.Route.Patch("/categories/:id", c.UpdateCategory)
	c.Route.Delete("/categories/:id", c.DeleteCategory)
}

// GetCategories godoc
//
//	@Summary		get All Categories
//	@Description	get All Categories
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	[]model.Category
//	@Router			/categories [get]
func (c Category) GetCategories(ctx *fiber.Ctx) error {

	results, err := c.Service.GetCategories(ctx.Context())
	if err != nil {
		return err

	}
	ctx.JSON(results)
	return nil
}

// AddCategories godoc
//
//	@Summary		Add New Category
//	@Description	Add New Category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@param			request	body		dto.CategoryDTO	true	"category data"
//	@Success		200		{object}	model.Category
//	@Router			/categories [Post]
func (c Category) AddCategory(ctx *fiber.Ctx) error {
	CategoryDTO := new(dto.CategoryDTO)
	_ = ctx.BodyParser(CategoryDTO)

	results, err := c.Service.AddCategory(ctx.Context(), CategoryDTO)
	if err != nil {
		return err

	}
	ctx.JSON(results)
	return nil
}

// UpdateCategories godoc
//
//	@Summary		Update New Category
//	@Description	Update New Category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@param			request		body		dto.CategoryDTO	true	"category data"
//	@param			category_id	path		int				true	"category id"
//	@Success		200			{object}	model.Category
//	@Router			/categories/{category_id} [Patch]
func (c Category) UpdateCategory(ctx *fiber.Ctx) error {
	CategoryDTO := new(dto.CategoryDTO)
	_ = ctx.BodyParser(CategoryDTO)

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	results, err := c.Service.UpdateCategory(ctx.Context(), id, CategoryDTO)

	if err != nil {
		return err

	}
	ctx.JSON(results)
	return nil
}

// DELETECategories godoc
//
//	@Summary		delete  Category
//	@Description	delete  Category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int		true	"Category ID"
//	@Success		200	{string}	string	"deleted successfully"
//
//	@Router			/categories/{id} [Delete]
func (c Category) DeleteCategory(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	err = c.Service.DeleteCategory(ctx.Context(), id)
	if err != nil {
		return err

	}

	ctx.Write([]byte("deleted successfully"))
	return nil
}
