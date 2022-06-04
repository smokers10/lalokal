package main

import (
	"lalokal/infrastructure/configuration"
	"lalokal/infrastructure/injector"
	"lalokal/infrastructure/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {
	// app configuration
	application_configuration := configuration.ReadConfiguration().Application

	//set view engine
	engine := html.New("./infrastructure/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main",
		// Prefork: true,
	})

	// static folder
	app.Static("/", "./infrastructure/public")

	if application_configuration.Mode == "development" {
		app.Use(recover.New())
	}

	// invoke injection
	injector := injector.Injector()

	// router
	routes.Router(app, injector)

	log.Fatal(app.Listen(application_configuration.Port))
}
