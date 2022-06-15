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

func (tc *topicController) DashboardTopicPage(c *fiber.Ctx) error {
	return c.Render("user/dashboard_topic/dashboard", nil)
}

func (tc *topicController) Store(c *fiber.Ctx) error {
	body := topic.Topic{}
	c.BodyParser(&body)
	body.UserId = c.Locals("id").(string)

	res := tc.topicService.Store(&body)

	return c.Status(res.Status).JSON(res)
}

func (tc *topicController) Update(c *fiber.Ctx) error {
	body := topic.Topic{}
	c.BodyParser(&body)
	body.UserId = c.Locals("id").(string)

	res := tc.topicService.Update(&body)

	return c.Status(res.Status).JSON(res)
}

func (tc *topicController) ReadAll(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	res := tc.topicService.ReadAll(id)

	return c.Status(res.Status).JSON(res)
}

func (tc *topicController) Detail(c *fiber.Ctx) error {
	res := tc.topicService.Detail(c.Params("id"))

	return c.Status(res.Status).JSON(res)
}
