package routes

import (
	"lalokal/controller"
	"lalokal/infrastructure/injector"
	"lalokal/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, solvent *injector.InjectorSolvent) {
	// middleware
	userMiddleware := middleware.UserMiddlewareHandler(injector.Injector())
	guestMiddleware := middleware.GuestHandler(injector.Injector())

	// main controller
	mainController := controller.MainController(solvent)

	// login
	authController := mainController.LoginController()
	app.Get("/", guestMiddleware, authController.LoginPage)
	app.Post("/login/submission", guestMiddleware, authController.LoginSubmission)
	app.Get("/logout", userMiddleware, authController.Logout)

	// register
	registrationController := mainController.RegistrationController()
	registrationPath := app.Group("/registration")
	registrationPath.Get("/step-1", guestMiddleware, registrationController.EmailVerificationRequestPage)
	registrationPath.Get("/step-2", guestMiddleware, registrationController.VerificatePage)
	registrationPath.Get("/step-3", guestMiddleware, registrationController.RegistrationPage)
	registrationPath.Post("/email-verification-request", guestMiddleware, registrationController.EmailVerificationRequestSubmission)
	registrationPath.Post("/verificate-submission", guestMiddleware, registrationController.VerificateSubmission)
	registrationPath.Post("/registration-submission", guestMiddleware, registrationController.RegistrationSubmission)

	// user - dashboard
	userPath := app.Group("/user", userMiddleware)

	// topic
	topicController := mainController.TopicController()
	topicPath := userPath.Group("/topic")
	topicPath.Get("/", topicController.TopicPage)

	// test
	testController := mainController.TestController()
	testPath := app.Group("/user", userMiddleware)
	testPath.Get("/est", testController.Protected)
}
