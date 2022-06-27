package controller

import (
	"lalokal/domain/blasting_session"
	"lalokal/domain/selected_tweet"
	service "lalokal/service/blasting_session"

	"github.com/gofiber/fiber/v2"
)

type blastingSessionController struct {
	blastingSessionService blasting_session.Service
}

func (mc *mainController) BlastingSessionController() blastingSessionController {
	bss := service.BlastingSessionService(
		&mc.solvent.Repository.BlastLogRepository,
		&mc.solvent.Repository.BlastingSessionRepository,
		&mc.solvent.Repository.TwitterAPITokenRepository,
		&mc.solvent.Repository.KeywordRepository,
		&mc.solvent.TwitterHTTP,
	)

	return blastingSessionController{blastingSessionService: bss}
}

func (bss *blastingSessionController) ManageSessionBlasstingPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/session_blasting", nil)
}

func (bss *blastingSessionController) SessionBlastingPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/blasting_control_page", nil)
}

func (bss *blastingSessionController) GetAllCount(c *fiber.Ctx) error {
	res := bss.blastingSessionService.Count(c.Params("topic_id"))

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) GetAll(c *fiber.Ctx) error {
	res := bss.blastingSessionService.ReadAll(c.Params("topic_id"))

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) GetDetail(c *fiber.Ctx) error {
	res := bss.blastingSessionService.Detail(c.Params("blasting_session_id"))

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) Store(c *fiber.Ctx) error {
	body := blasting_session.BlastingSession{}
	c.BodyParser(&body)

	res := bss.blastingSessionService.Store(&body)

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) Update(c *fiber.Ctx) error {
	body := blasting_session.BlastingSession{}
	c.BodyParser(&body)

	res := bss.blastingSessionService.Update(&body)

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) Scrape(c *fiber.Ctx) error {
	bssID := c.Params("blasting_session_id")

	res := bss.blastingSessionService.Scrape(bssID)

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) Blast(c *fiber.Ctx) error {
	type data struct {
		Selected           []selected_tweet.SelectedTweet `json:"selected"`
		BlastringSessionId string                         `json:"blasting_session_id"`
	}
	body := data{}
	c.BodyParser(&body)

	res := bss.blastingSessionService.Blast(body.BlastringSessionId, body.Selected)

	return c.Status(res.Status).JSON(res)
}

func (bss *blastingSessionController) Monitoring(c *fiber.Ctx) error {
	res := bss.blastingSessionService.Monitoring(c.Params("blasting_session_id"))

	return c.Status(res.Status).JSON(res)
}
