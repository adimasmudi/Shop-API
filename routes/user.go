package routes

import (
	"Shop-API/controllers"
	"Shop-API/middleware"
	"Shop-API/repository"
	"Shop-API/services"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func UserRoute(api fiber.Router, db *gorm.DB) {
	userRepository := repository.NewRepository(db)
	userService := services.NewService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	// routes
	// user
	api.Post("/register",userHandler.Register)
	api.Post("/login", userHandler.Login)
	api.Get("/profile", middleware.Auth(), userHandler.GetProfile)
	api.Put("/profile/update", middleware.Auth(), userHandler.UpdateProfile)
}