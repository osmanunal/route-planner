package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"route-planner/api/internal/handler"
	"route-planner/service"
)

func Setup(app *fiber.App, db *bun.DB) {
	locationService := service.NewLocationService(db)

	locationHandler := handler.NewLocationHandler(locationService)

	api := app.Group("/api")

	api.Get("/location", locationHandler.GetAll)
	api.Get("/location/:id", locationHandler.GetByID)
	api.Post("/location", locationHandler.Create)
	api.Put("/location/:id", locationHandler.Update)
	api.Delete("/location/:id", locationHandler.Delete)

}
