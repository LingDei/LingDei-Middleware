package main

import (
	"LingDei-Middleware/config"

	_ "LingDei-Middleware/docs"

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
	video.Post("/update", handler.UpdateVideoHandler)
	video.Get("/list", handler.GetVideoListHandler)
	video.Get("/get", handler.GetVideoHandler)
	video.Delete("/delete", handler.DeleteVideoHandler)

	// Category
	category := app.Group("/category")
	category.Post("/add", handler.AddCategoryHandler)
	category.Get("/list", handler.GetCategoryListHandler)
	category.Get("/get", handler.GetCategoryHandler)
	category.Delete("/delete", handler.DeleteCategoryHandler)
	category.Post("/update", handler.UpdateCategoryHandler)
}
