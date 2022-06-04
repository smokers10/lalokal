package controller

import (
	"lalokal/domain/user"
	"lalokal/domain/verification"
	service "lalokal/service/user"

	"github.com/gofiber/fiber/v2"
)

type registrationController struct {
	userService user.Service
}

func (mc *mainController) RegistrationController() *registrationController {
	userService := service.UserService(&mc.solvent.Repository.UserRepository, &mc.solvent.Repository.ForgotPasswordRepository,
		&mc.solvent.Repository.VerificationRepository, &mc.solvent.Encryption, &mc.solvent.JsonWebToken,
		&mc.solvent.Identifier, &mc.solvent.Mailer)

	return &registrationController{userService: userService}
}

func (rc *registrationController) EmailVerificationRequestPage(c *fiber.Ctx) error {
	return c.Render("registration/step-1-verification-request", nil)
}

func (rc *registrationController) VerificatePage(c *fiber.Ctx) error {
	return c.Render("registration/step-2-verificate", nil)
}

func (rc *registrationController) RegistrationPage(c *fiber.Ctx) error {
	return c.Render("registration/step-3-registration", nil)
}

func (rc *registrationController) EmailVerificationRequestSubmission(c *fiber.Ctx) error {
	email := c.FormValue("email")

	res := rc.userService.VerificationRequest(email)

	return c.Status(res.Status).JSON(res)
}

func (rc *registrationController) VerificateSubmission(c *fiber.Ctx) error {
	body := verification.Verification{}
	c.BodyParser(&body)

	res := rc.userService.VerificateEmail(&body)

	return c.Status(res.Status).JSON(res)
}

func (rc *registrationController) RegistrationSubmission(c *fiber.Ctx) error {
	body := user.RegisterData{}
	c.BodyParser(&body)

	res := rc.userService.Register(&body)

	return c.Status(res.Status).JSON(res)
}
