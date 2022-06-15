package controller

import (
	"lalokal/domain/blasting_session"
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
		&mc.solvent.Repository.SelectedTweetRepository,
	)

	return blastingSessionController{blastingSessionService: bss}
}

func (bss *blastingSessionController) ManageSessionBlasstingPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/session_blasting", nil)
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
