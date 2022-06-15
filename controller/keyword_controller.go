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
