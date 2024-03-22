package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

func (h Handler) GetInstance(c *fiber.Ctx) error {
	id := c.Params("id")
	convertedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		h.logger.Error(
			err.Error(),
			zap.String("arch level", "handler"),
			zap.String("handler", "GetInstance"),
		)
		return c.SendStatus(400)
	}

	if convertedId <= 0 {
		h.logger.Error(
			"invalid body",
			zap.String("arch level", "handler"),
			zap.String("handler", "GetInstance"),
		)
		return c.SendStatus(400)
	}

	instances, err := h.services.GetInstance(convertedId)
	if err != nil {
		h.logger.Error(
			err.Error(),
			zap.String("arch level", "handler"),
			zap.String("handler", "GetInstance"),
		)
		return c.SendStatus(500)
	}

	if len(instances) == 0 {
		h.logger.Error(
			"this instance doesn't exist",
			zap.String("arch level", "handler"),
			zap.String("handler", "GetInstance"),
		)
		return c.SendStatus(400)
	}

	jsonInstancesData, err := json.Marshal(instances)
	if err != nil {
		h.logger.Error(
			err.Error(),
			zap.String("arch level", "handler"),
			zap.String("handler", "GetInstance"),
		)
		return c.SendStatus(500)
	}

	return c.Send(jsonInstancesData)
}

func (h Handler) GetAllInstances(c *fiber.Ctx) error {
	// TODO: implement this
	return nil
}

func (h Handler) CreateInstance(c *fiber.Ctx) error {
	// TODO: implement this
	return nil
}

func (h Handler) UpdateInstance(c *fiber.Ctx) error {
	// TODO: implement this
	return nil
}

func (h Handler) DeleteInstance(c *fiber.Ctx) error {
	// TODO: implement this
	return nil
}
