package main

import (
	"Shop-API/helper"
	"Shop-API/routes"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main(){

	// get data from database
	dsn := helper.EnvDBURL()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	
	app := fiber.New()

	app.Use(middleware.Logger())

	

	api := app.Group("/api/v1")

	// routes
	routes.UserRoute(api, db)
	routes.TokoRouter(api, db)
	

	app.Listen(":6000")


}