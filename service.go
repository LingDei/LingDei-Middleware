package main

import (
	"LingDei-Middleware/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/swagger"

	handler "LingDei-Middleware/handler"
)

var jwtConfig = jwtware.New(jwtware.Config{
	SigningMethod: "RS256",
	SigningKey:    handler.PrivateKey.Public(),
	ErrorHandler:  config.CustomJwtError(),
})

func regiserService(app *fiber.App) {
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 以下为权限控制接口
	app.Use(jwtConfig)

	// Video
	video := app.Group("/video")
	video.Post("/add", handler.AddVideoHandler)
	video.Get("/list", handler.GetVideoListHandler)
	video.Get("/get", handler.GetVideoHandler)
	video.Delete("/delete", handler.DeleteVideoHandler)
	video.Post("/update", handler.UpdateVideoHandler)
}
