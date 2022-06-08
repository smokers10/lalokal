package controller

import (
	"lalokal/domain/topic"
	service "lalokal/service/topic"

	"github.com/gofiber/fiber/v2"
)

type topicController struct {
	topicService topic.Service
}

func (mc *mainController) TopicController() *topicController {
	ts := service.TopicService(&mc.solvent.Repository.TopicRepository)
	return &topicController{topicService: ts}
}

func (tc *topicController) TopicPage(c *fiber.Ctx) error {
	return c.Render("user/topic/topic", nil)
}
