package handlers

import (
    "github.com/gofiber/fiber/v2"
    "social_network/utils"
)

type UserHandler struct {
    Storage *utils.AuthStorage
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
    // Логика получения профиля
    return c.SendString("Profile endpoint")
}
