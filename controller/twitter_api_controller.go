package controller

import (
	"lalokal/domain/twitter_api_token"
	service "lalokal/service/twitter_api_token"

	"github.com/gofiber/fiber/v2"
)

type twitterAPIController struct {
	twitterAPIService twitter_api_token.Service
}

func (mc *mainController) TwitterAPIController() twitterAPIController {
	tac := service.TwitterAPIService(&mc.solvent.Repository.TwitterAPITokenRepository)

	return twitterAPIController{twitterAPIService: tac}
}

func (tac *twitterAPIController) ManageTwitterAPIPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/api_key", nil)
}

func (tac *twitterAPIController) Store(c *fiber.Ctx) error {
	body := twitter_api_token.TwitterAPIToken{}
	c.BodyParser(&body)

	res := tac.twitterAPIService.Store(&body)

	return c.Status(res.Status).JSON(res)
}

func (tac *twitterAPIController) Read(c *fiber.Ctx) error {
	res := tac.twitterAPIService.Read(c.Params("topic_id"))

	return c.Status(res.Status).JSON(res)
}
