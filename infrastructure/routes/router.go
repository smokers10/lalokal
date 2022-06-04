package routes

import (
	"lalokal/controller"
	"lalokal/infrastructure/injector"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, solvent *injector.InjectorSolvent) {
	mainController := controller.MainController(solvent)

	authController := mainController.LoginController()
	app.Get("/", authController.LoginPage)

	// register
	registrationController := mainController.RegistrationController()
	registrationPath := app.Group("/registration")
	registrationPath.Get("/step-1", registrationController.EmailVerificationRequestPage)
	registrationPath.Get("/step-2", registrationController.VerificatePage)
	registrationPath.Get("/step-3", registrationController.RegistrationPage)
	registrationPath.Post("/email-verification-request", registrationController.EmailVerificationRequestSubmission)
	registrationPath.Post("/verificate-submission", registrationController.VerificateSubmission)
	registrationPath.Post("/registration-submission", registrationController.RegistrationSubmission)

	// test
	testController := mainController.TestController()
	testPath := app.Group("/test")
	testPath.Get("/", testController.Protected)
}
