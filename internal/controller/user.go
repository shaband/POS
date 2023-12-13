package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/shaband/POS/internal/helper"
	"github.com/shaband/POS/internal/model"
	"github.com/shaband/POS/internal/model/dto"
	"github.com/shaband/POS/internal/output"
	"github.com/shaband/POS/internal/service"
	"go.uber.org/fx"
)

type User struct {
	fx.In
	Encryptor *helper.Encryptor
	Validator *helper.Validator
	Service   *service.User
	Route     fiber.Router `name:"api-v1"`
}

func RegisterUser(c User) {
	c.Route.Get("/users", c.GetUsers)
	c.Route.Post("/users", c.AddUser)
	c.Route.Patch("/users/:id", c.UpdateUser)
	c.Route.Delete("/users/:id", c.DeleteUser)
}

// GetUsers godoc
//
//	@Summary		get All Users
//	@Description	get All Users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	output.User
//	@Router			/users [get]
func (c *User) GetUsers(ctx *fiber.Ctx) error {
	users, err := c.Service.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	result := []*output.User{}

	for _, user := range users {
		result = append(result, &output.User{
			Id:    uint(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		})
	}

	return ctx.JSON(result)
}

// AddCategories godoc
//
//	@Summary		Add New User
//	@Description	Add New User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@param			request	body		dto.CreateUserDTO	true	"user data"
//	@Success		200		{object}	output.User
//
// @Failure      	400  	{object}  	output.HTTPError
// @Router			/users [Post]
func (c User) AddUser(ctx *fiber.Ctx) error {
	UserDTO := new(dto.CreateUserDTO)
	_ = ctx.BodyParser(UserDTO)
	Matched := c.Validator.StringMatches(UserDTO.Password, UserDTO.PasswordConfirmation)
	if !Matched {
		return errors.New("Password doesn't match password confirmation")
	}
	hashed, _ := c.Encryptor.HashPassword(UserDTO.Password)
	UserDTO.Password = hashed
	isValid, err := c.Validator.Validate(UserDTO)
	if isValid {
		results, err := c.Service.AddUser(ctx.Context(), UserDTO)
		if err == nil {
			return ctx.JSON(output.User{
				Id:   uint(results.ID),
				Name: results.Name,
			})
		}
	}
	return err
}

// UpdateCategories godoc
//
//	@Summary		Update  User
//	@Description	Update  User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@param			request		body		dto.UpdateUserDTO	true	"user data"
//	@param			user_id	path		int				true	"user id"
//
//	@Success		200			{object}	model.User
//
// @Failure      	400  		{object}  	output.HTTPError
// @Router			/users/{user_id} [Patch]
func (c User) UpdateUser(ctx *fiber.Ctx) error {
	UserDTO := new(dto.UpdateUserDTO)
	_ = ctx.BodyParser(UserDTO)
	if UserDTO.Password != "" {
		Matched := c.Validator.StringMatches(UserDTO.Password, UserDTO.PasswordConfirmation)
		if !Matched {
			return errors.New("Password doesn't match password confirmation")
		}
		hashed, _ := c.Encryptor.HashPassword(UserDTO.Password)
		UserDTO.Password = hashed
	}
	isValid, err := c.Validator.Validate(UserDTO)
	if isValid {
		id, err := ctx.ParamsInt("id")
		if err == nil {
			var results *model.User
			results, err = c.Service.UpdateUser(ctx.Context(), id, UserDTO)
			if err != nil {
				return err
			}
			return ctx.JSON(output.User{
				Id:   uint(results.ID),
				Name: results.Name,
			})
		}
	}
	return err
}

// DeleteUser godoc
//
//	@Summary		delete  User
//	@Description	delete  User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		int		true	"User ID"
//	@Success		200	{string}	string	"deleted successfully"
//
//	@Router			/users/{id} [Delete]
func (c User) DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err == nil {
		err = c.Service.DeleteUser(ctx.Context(), id)
		if err == nil {
			_, err = ctx.WriteString("deleted successfully")
		}
	}

	return err
}
