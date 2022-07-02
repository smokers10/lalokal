package controller

import (
	"lalokal/domain/user"
	service "lalokal/service/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type forgotPasswordController struct {
	userService user.Service
	session     *session.Store
}

func (mc *mainController) ForgotPasswordController() forgotPasswordController {
	userService := service.UserService(&mc.solvent.Repository.UserRepository, &mc.solvent.Repository.ForgotPasswordRepository,
		&mc.solvent.Repository.VerificationRepository, &mc.solvent.Encryption, &mc.solvent.JsonWebToken,
		&mc.solvent.Identifier, &mc.solvent.Mailer)

	return forgotPasswordController{userService: userService, session: &mc.solvent.Session}
}

func (fpc *forgotPasswordController) ForgotPasswordPage(c *fiber.Ctx) error {
	return c.Render("forgot-password/forgot-password", nil)
}
