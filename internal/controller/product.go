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

type Product struct {
	fx.In
	Uploader  *helper.Uploader
	Validator *helper.Validator
	Service   *service.Product
	Route     fiber.Router `name:"api-v1"`
}

func RegisterProduct(c Product) {
	c.Route.Get("/products", c.GetProducts)
	c.Route.Post("/products", c.AddProduct)
	c.Route.Patch("/products/:id", c.UpdateProduct)
	c.Route.Delete("/products/:id", c.DeleteProduct)
}

// GetProducts godoc
//
//	@Summary		get All Products
//	@Description	get All Products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	output.Product
//	@Router			/products [get]
func (c *Product) GetProducts(ctx *fiber.Ctx) error {
	products, err := c.Service.GetProducts(ctx.Context())
	if err != nil {
		return err
	}
	result := []*output.Product{}

	for _, product := range products {
		result = append(result, &output.Product{
			ID:        product.ID,
			Name:      product.Name,
			Code:      product.Code,
			CostPrice: product.CostPrice,
			SellPrice: product.SellPrice,
			Image:     product.Image,
		})
	}

	return ctx.JSON(result)
}

// AddCategories godoc
//
//	@Summary		Add New Product
//	@Description	Add New Product
//	@Tags			products
//	@Accept			mpfd
//	@Produce		mpfd
//	@param			name		formData	string	true	"enter product name"
//	@param			code		formData	string	true	"enter product code"
//	@param			cost_price	formData	number	true	"enter product cost price"
//	@param			sell_price	formData	number	true	"enter product sell price"
//	@param			category_id	formData	int		true	"category ID"
//	@param			image		formData	file	false	"image"
//	@Success		200			{object}	output.Product
//
//	@Failure		400			{object}	output.HTTPError
//	@Router			/products [Post]
func (c Product) AddProduct(ctx *fiber.Ctx) error {
	ProductDTO := new(dto.ProductDTO)
	_ = ctx.BodyParser(ProductDTO)
	path, err := c.Uploader.SaveFIle(ctx, "image", "products")
	if err == nil {
		ProductDTO.Image = path
	}
	isValid, err := c.Validator.Validate(ProductDTO)
	if isValid {
		results, err := c.Service.AddProduct(ctx.Context(), ProductDTO)
		if err == nil {
			return ctx.JSON(output.Product{
				ID:        results.ID,
				Name:      results.Name,
				Code:      results.Code,
				CostPrice: results.CostPrice,
				SellPrice: results.SellPrice,
				Image:     results.Image,
			})
		}
	}
	return err
}

// UpdateCategories godoc
//
//	@Summary		Update  Product
//	@Description	Update  Product
//	@Tags			products
//	@Accept			mpfd
//	@Produce		mpfd
//	@param			name		formData	string	true	"enter product name"
//	@param			code		formData	string	true	"enter product code"
//	@param			cost_price	formData	number	true	"enter product cost price"
//	@param			sell_price	formData	number	true	"enter product sell price"
//	@param			category_id	formData	int		true	"category ID"
//	@param			image		formData	file	true	"image"
//	@param			product_id	path		int		true	"Product ID"
//
//	@Success		200			{object}	model.Product
//
//	@Failure		400			{object}	output.HTTPError
//	@Router			/products/{product_id} [Patch]
func (c Product) UpdateProduct(ctx *fiber.Ctx) error {
	ProductDTO := new(dto.ProductDTO)
	_ = ctx.BodyParser(ProductDTO)

	path, err := c.Uploader.SaveFIle(ctx, "image", "products")
	if err == nil {
		ProductDTO.Image = path
	}
	isValid, err := c.Validator.Validate(ProductDTO)
	if isValid {
		id, err := ctx.ParamsInt("id")
		if err == nil {
			var results *model.Product
			results, err = c.Service.UpdateProduct(ctx.Context(), id, ProductDTO)
			if err != nil {
				return err
			}
			return ctx.JSON(output.Product{
				ID:        results.ID,
				Name:      results.Name,
				Code:      results.Code,
				CostPrice: results.CostPrice,
				SellPrice: results.SellPrice,
				Image:     results.Image,
			})
		}
	}
	return err
}

// DeleteProduct godoc
//
//	@Summary		delete  Product
//	@Description	delete  Product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int		true	"Product ID"
//	@Success		200	{string}	string	"deleted successfully"
//
//	@Router			/products/{id} [Delete]
func (c Product) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err == nil {
		err = c.Service.DeleteProduct(ctx.Context(), id)
		if err == nil {
			_, err = ctx.WriteString("deleted successfully")
		}
	}

	return err
}
