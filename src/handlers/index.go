package handlers

import (
	auth "liftlab/src/handlers/Auth"
	users "liftlab/src/handlers/Users"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}
type UserHandler struct{}

var Auth = new(AuthHandler)
var User = new(UserHandler)

func (a *AuthHandler) Login(c *fiber.Ctx) error {
	return auth.Login(c)
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	return users.GetUser(c)
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return users.UpdateUser(c)
}
