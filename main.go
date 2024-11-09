package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
)

var runningInDocker bool

func main() {
	runningInDocker = os.Getenv("RunInDocker") != ""

	// Getting Application port from the Environment
	port, portExists := os.LookupEnv("PORT")
	if !portExists {
		port = "9000"
	}

	// Getting LocalIP (Only for Docker Image)
	ipAddress, err := getLocalIP()
	if err != nil {
		log.Println("Warning: " + err.Error())
	}

	// Init & Configure Fiber
	app := fiber.New()
	configureApp(app)

	log.Fatal(app.Listen(ipAddress + ":" + port))
}
