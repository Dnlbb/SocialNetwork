package main

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
    "github.com/sirupsen/logrus"
    "social_network/handlers"
    "social_network/utils"
)

const (
    contextKeyUser = "user"
)

func main() {
    app := fiber.New()

    authStorage := &utils.AuthStorage{Users: map[string]utils.User{}}
    authHandler := &handlers.AuthHandler{Storage: authStorage}
    userHandler := &handlers.UserHandler{Storage: authStorage}

    // Группа обработчиков, которые доступны неавторизованным пользователям
    publicGroup := app.Group("")
    publicGroup.Post("/register", authHandler.Register)
    publicGroup.Post("/login", authHandler.Login)

    // Группа обработчиков, которые требуют авторизации
    authorizedGroup := app.Group("")
    authorizedGroup.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{
            Key: utils.JwtSecretKey,
        },
        ContextKey: contextKeyUser,
    }))
    authorizedGroup.Get("/profile", userHandler.Profile)

    logrus.Fatal(app.Listen(":80"))
}
