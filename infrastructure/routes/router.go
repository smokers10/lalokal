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
	topicPath := userPath.Group("/topic", userMiddleware)
	topicPath.Get("/", topicController.TopicPage)
	topicPath.Get("/get", topicController.ReadAll)
	topicPath.Get("/get/:id", topicController.Detail)
	topicPath.Post("/store", topicController.Store)
	topicPath.Post("/update", topicController.Update)

	// dashboard
	dashboardPath := topicPath.Group("/dashboard")
	dashboardPath.Get("/", topicController.DashboardTopicPage)

	// blasting session
	blastingSessionPath := dashboardPath.Group("/blasting-session")
	blastingSessionController := mainController.BlastingSessionController()
	blastingSessionPath.Get("/", blastingSessionController.ManageSessionBlasstingPage)
	blastingSessionPath.Get("/get-count/:topic_id", blastingSessionController.GetAllCount)
	blastingSessionPath.Get("/get-all/:topic_id", blastingSessionController.GetAll)
	blastingSessionPath.Get("/detail/:blasting_session_id", blastingSessionController.GetDetail)
	blastingSessionPath.Post("/store", blastingSessionController.Store)
	blastingSessionPath.Post("/update", blastingSessionController.Update)

	// wtitter api key
	twitterAPIPath := dashboardPath.Group("/twitter-api")
	twitterAPIController := mainController.TwitterAPIController()
	twitterAPIPath.Get("/", twitterAPIController.ManageTwitterAPIPage)
	twitterAPIPath.Get("/get/:topic_id", twitterAPIController.Read)
	twitterAPIPath.Post("/store", twitterAPIController.Store)

	// keyword
	keywordPath := dashboardPath.Group("/keyword")
	keywordController := mainController.KeywordController()
	keywordPath.Get("/", keywordController.ManageKeywordPage)

	// test
	testController := mainController.TestController()
	testPath := app.Group("/user", userMiddleware)
	testPath.Get("/est", testController.Protected)
}
