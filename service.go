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

	// 注册路由
	category := app.Group("/category")
	video := app.Group("/video")
	like := app.Group("/like")
	collect := app.Group("/collect")

	// Category
	category.Get("/list", handler.GetCategoryListHandler)
	category.Get("/get", handler.GetCategoryHandler)

	// Video
	video.Get("/list", handler.GetVideoListHandler)
	video.Get("/get", handler.GetVideoHandler)

	// Like
	like.Get("/count", handler.GetLikeCountHandler)

	// 以下为权限控制接口
	app.Use(jwtConfig)

	// Video
	video.Post("/add", handler.AddVideoHandler)
	video.Post("/update", handler.UpdateVideoHandler)
	video.Delete("/delete", handler.DeleteVideoHandler)

	// Category
	category.Post("/add", handler.AddCategoryHandler)
	category.Post("/update", handler.UpdateCategoryHandler)
	category.Delete("/delete", handler.DeleteCategoryHandler)

	// Like
	like.Post("/add", handler.AddLikeHandler)
	like.Get("/list", handler.GetLikeListHandler)
	like.Get("/check", handler.CheckLikeExistHandler)
	like.Delete("/delete", handler.DeleteLikeHandler)

	// Collect
	collect.Post("/add", handler.AddCollectHandler)
	collect.Get("/list", handler.GetCollectListHandler)
	collect.Get("/check", handler.CheckCollectExistHandler)
	collect.Delete("/delete", handler.DeleteCollectHandler)
}
