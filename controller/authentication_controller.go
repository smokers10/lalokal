package controller

import (
	"lalokal/domain/user"
	service "lalokal/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type authenticationController struct {
	userService user.Service
	session     *session.Store
}

func (mc *mainController) LoginController() authenticationController {
	userService := service.UserService(&mc.solvent.Repository.UserRepository, &mc.solvent.Repository.ForgotPasswordRepository,
		&mc.solvent.Repository.VerificationRepository, &mc.solvent.Encryption, &mc.solvent.JsonWebToken,
		&mc.solvent.Identifier, &mc.solvent.Mailer)

	return authenticationController{userService: userService, session: &mc.solvent.Session}
}

func (ac *authenticationController) LoginPage(c *fiber.Ctx) error {
	return c.Render("authentication/login", nil)
}

func (ac *authenticationController) LoginSubmission(c *fiber.Ctx) error {
	body := user.LoginData{}
	c.BodyParser(&body)

	res := ac.userService.Login(&body)

	if res.Success {
		sess, err := ac.session.Get(c)
		if err != nil {
			panic(err)
		}

		sess.Set("token", res.Token)
		sess.Save()
	}

	return c.Status(res.Status).JSON(res)
}

func (ac *authenticationController) Logout(c *fiber.Ctx) error {
	sess, err := ac.session.Get(c)
	if err != nil {
		panic(err)
	}

	sess.Destroy()

	return c.Redirect("/")
}
