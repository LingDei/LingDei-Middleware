package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "LingDei-Middleware/docs"
)

//	@title						灵嘚中间件
//	@version					v1.0.0
//	@description				灵嘚中间件（LingDei-Middleware）
//	@host						127.0.0.1:9000
//	@BasePath					/
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	// 创建 Fiber 实例
	app := fiber.New()

	// 启动 CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// 注册服务
	// regiserService(app)

	// 启动服务
	app.Listen(":9000")
}
