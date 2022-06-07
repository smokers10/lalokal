package routes

import (
	"lalokal/controller"
	"lalokal/infrastructure/injector"
	"lalokal/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, solvent *injector.InjectorSolvent) {
	mainController := controller.MainController(solvent)

	// login
	authController := mainController.LoginController()
	app.Get("/", authController.LoginPage)
	app.Post("/login/submission", authController.LoginSubmission)
	app.Get("/logout", middleware.UserMiddleware(injector.Injector()), authController.Logout)

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
	testPath := app.Group("/test", middleware.UserMiddleware(injector.Injector()))
	testPath.Get("/", testController.Protected)
}
