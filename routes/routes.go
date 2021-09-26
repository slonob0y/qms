package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slonob0y/qms/handler"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)
	app.Get("/api/user", handler.User)
	app.Post("/api/logout", handler.Logout)

}