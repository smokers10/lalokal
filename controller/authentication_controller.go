package controller

import "github.com/gofiber/fiber/v2"

type authenticationController struct{}

func (mc *mainController) LoginController() authenticationController {
	return authenticationController{}
}

func (ac *authenticationController) LoginPage(c *fiber.Ctx) error {
	return c.Render("authentication/login", nil)
}
