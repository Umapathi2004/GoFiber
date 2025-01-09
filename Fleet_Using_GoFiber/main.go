package main

import (
	"Fleet_GoFiber/Routes"

	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()

	Routes.DriverRoute(app.Group("/driver"))
	Routes.VehicleRoute(app.Group("/vehicle"))

	app.Listen(fmt.Sprintf(":%v", PORT))
}
