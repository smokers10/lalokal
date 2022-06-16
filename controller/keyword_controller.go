package controller

import (
	"lalokal/domain/keyword"
	service "lalokal/service/keyword"

	"github.com/gofiber/fiber/v2"
)

type keywordController struct {
	keywordService keyword.Service
}

func (mc *mainController) KeywordController() keywordController {
	ks := service.KeywordService(&mc.solvent.Repository.KeywordRepository)
	return keywordController{keywordService: ks}
}

func (kc *keywordController) ManageKeywordPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/keyword", nil)
}

func (kc *keywordController) Store(c *fiber.Ctx) error {
	body := keyword.Keyword{}
	c.BodyParser(&body)

	res := kc.keywordService.Store(&body)

	return c.Status(res.Status).JSON(res)
}

func (kc *keywordController) Delete(c *fiber.Ctx) error {
	body := keyword.Keyword{}
	c.BodyParser(&body)

	res := kc.keywordService.Delete(body.Id)

	return c.Status(res.Status).JSON(res)
}

func (kc *keywordController) GetAll(c *fiber.Ctx) error {
	id := c.Params("id")

	res := kc.keywordService.ReadAll(id)

	return c.Status(res.Status).JSON(res)
}
