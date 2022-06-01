package routes

import (
	"lalokal/controller"
	"lalokal/infrastructure/injector"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, solvent *injector.InjectorSolvent) {
	mainController := controller.MainController(solvent)

	authController := mainController.LoginController()
	app.Get("/login", authController.LoginPage)
}
