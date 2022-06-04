package controller

import "github.com/gofiber/fiber/v2"

type testController struct{}

func (mc *mainController) TestController() *testController {
	return &testController{}
}

func (tc *testController) Protected(c *fiber.Ctx) error {
	return c.Render("protected", nil)
}
