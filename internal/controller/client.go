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

type Client struct {
	fx.In
	Encryptor *helper.Encryptor
	Validator *helper.Validator
	Service   *service.Client
	Route     fiber.Router `name:"api-v1"`
}

func RegisterClient(c Client) {
	c.Route.Get("/clients", c.GetClients)
	c.Route.Post("/clients", c.AddClient)
	c.Route.Patch("/clients/:id", c.UpdateClient)
	c.Route.Delete("/clients/:id", c.DeleteClient)
}

// GetClients godoc
//
//	@Summary		get All Clients
//	@Description	get All Clients
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	output.Client
//	@Router			/clients [get]
func (c *Client) GetClients(ctx *fiber.Ctx) error {
	clients, err := c.Service.GetClients(ctx.Context())
	if err != nil {
		return err
	}
	result := []*output.Client{}

	for _, client := range clients {
		result = append(result, &output.Client{
			Id:    uint(client.ID),
			Name:  client.Name,
			Email: client.Email,
			Phone: client.Phone,
		})
	}

	return ctx.JSON(result)
}

// AddCategories godoc
//
//	@Summary		Add New Client
//	@Description	Add New Client
//	@Tags			clients
//	@Accept			mpfd
//	@Produce		mpfd
//	@param			name	    formData		string	true	"enter client name"
//	@param			email	    formData		string	true	"enter client email"
//	@param			phone	    formData		number	true	"enter client phone"
//	@param			is_client	formData		boolean	false	" is client ?"
//	@param			is_supplier	formData		boolean	false	"is supplier ?"
//	@param			image    formData	file		        true	"image"
//	@Success		200		{object}	output.Client
//
// @Failure      	400  	{object}  	output.HTTPError
// @Router			/clients [Post]
func (c Client) AddClient(ctx *fiber.Ctx) error {
	ClientDTO := new(dto.CreateClientDTO)
	_ = ctx.BodyParser(ClientDTO)
	isValid, err := c.Validator.Validate(ClientDTO)
	if isValid {
		results, err := c.Service.AddClient(ctx.Context(), ClientDTO)
		if err == nil {
			return ctx.JSON(output.Client{
				Id:   uint(results.ID),
				Name: results.Name,
			})
		}
	}
	return err
}

// UpdateCategories godoc
//
//	@Summary		Update  Client
//	@Description	Update  Client
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@param			request		body		dto.UpdateClientDTO	true	"client data"
//	@param			client_id	path		int				true	"client id"
//
//	@Success		200			{object}	model.Client
//
// @Failure      	400  		{object}  	output.HTTPError
// @Router			/clients/{client_id} [Patch]
func (c Client) UpdateClient(ctx *fiber.Ctx) error {
	ClientDTO := new(dto.UpdateClientDTO)
	_ = ctx.BodyParser(ClientDTO)

	isValid, err := c.Validator.Validate(ClientDTO)
	if isValid {
		id, err := ctx.ParamsInt("id")
		if err == nil {
			var results *model.Client
			results, err = c.Service.UpdateClient(ctx.Context(), id, ClientDTO)
			if err != nil {
				return err
			}
			return ctx.JSON(output.Client{
				Id:   uint(results.ID),
				Name: results.Name,
			})
		}
	}
	return err
}

// DeleteClient godoc
//
//	@Summary		delete  Client
//	@Description	delete  Client
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int		true	"Client ID"
//	@Success		200	{string}	string	"deleted successfully"
//
//	@Router			/clients/{id} [Delete]
func (c Client) DeleteClient(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err == nil {
		err = c.Service.DeleteClient(ctx.Context(), id)
		if err == nil {
			_, err = ctx.WriteString("deleted successfully")
		}
	}

	return err
}
