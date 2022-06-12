package middleware

import (
	"lalokal/infrastructure/injector"
	"lalokal/infrastructure/middleware/user_middleware"

	"github.com/gofiber/fiber/v2"
)

func UserMiddlewareHandler(injector *injector.InjectorSolvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := injector.Session.Get(c)
		if err != nil {
			panic(err)
		}

		token := sess.Get("token")

		if token == nil {
			return c.Redirect("/")
		}

		middlewareProc := user_middleware.UserMiddleware(
			&injector.Repository.UserRepository,
			&injector.Repository.VerificationRepository,
			&injector.JsonWebToken,
		)

		res := middlewareProc.Process(token.(string))

		if !res.Is_pass && res.Reason == "empty token" {
			return c.Redirect("/")
		}

		if !res.Is_pass && res.Reason == "error" {
			panic(res.Reason)
		}

		if !res.Is_pass && res.Reason == "unregistered" {
			return c.Redirect("/registration/step-1")
		}

		if !res.Is_pass && res.Reason == "not verified" {
			return c.Redirect("/registration/step-1")
		}

		c.Locals("id", res.Claim.Id)
		c.Locals("email", res.Claim.Email)

		return c.Next()
	}
}

func GuestHandler(injector *injector.InjectorSolvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := injector.Session.Get(c)
		if err != nil {
			panic(err)
		}

		token := sess.Get("token")

		if token != nil {
			return c.Redirect("/user/topic")
		}

		return c.Next()
	}
}
