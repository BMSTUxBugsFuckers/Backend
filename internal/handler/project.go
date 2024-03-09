package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h Handler) GetProject(c *fiber.Ctx) error {
	strId := c.Params("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	project, err := h.services.GetProject(id)
	if err != nil {
		return err
	}

	jsonProject, err := json.Marshal(project)
	if err != nil {
		return err
	}

	return c.Send(jsonProject)
}
