package handlers

import (
    "github.com/gofiber/fiber/v2"
    "social_network/utils"
)

type AuthHandler struct {
    Storage *utils.AuthStorage
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    // Логика регистрации
    return c.SendString("Register endpoint")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    // Логика логина
    return c.SendString("Login endpoint")
}
