package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// helloHandler responds to requests with a greeting using Fiber.
func helloHandler(c *fiber.Ctx) error {
	// Get the hostname to show which container is handling the request.
	hostname, err := os.Hostname()
	if err != nil {
		// If we can't get the hostname, log the error and send a generic message.
		log.Printf("Could not get hostname: %v", err)
		hostname = "unknown"
	}

	// Prepare the message and log the request.
	message := fmt.Sprintf("Hello, World! from container: %s\n", hostname)
	log.Printf("Received request from %s on host %s", c.IP(), hostname)

	// Send the string response to the client.
	return c.SendString(message)
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Create a new Fiber instance.
	app := fiber.New()

	// Use the logger middleware to log HTTP requests.
	app.Use(logger.New())

	// Set the port for the server to listen on. Default to 8080.
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Register the handler for the GET "/" route.
	app.Get("/", helloHandler)

	// Start the Fiber server.
	log.Printf("Server starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
