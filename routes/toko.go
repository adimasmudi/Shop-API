package routes

import (
	"Shop-API/controllers"
	"Shop-API/middleware"
	"Shop-API/repository"
	"Shop-API/services"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func TokoRouter(api fiber.Router, db *gorm.DB){
	tokoRepository := repository.NewRepositoryToko(db)
	tokoService := services.NewServiceToko(tokoRepository)
	tokoHandler := controllers.NewTokoHandler(tokoService)

	// routes
	api.Get("/toko",middleware.Auth(), tokoHandler.GetMyToko)
	api.Get("/toko/all",tokoHandler.GetAllToko)
	api.Put("/toko/:id_toko",middleware.Auth(),tokoHandler.UpdateToko)
}