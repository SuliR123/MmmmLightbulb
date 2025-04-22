package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
    "github.com/SuliR123/MmmmLightbulb/internal/handlers"
)

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Define your routes
    app.Get("/hello", handlers.GetHello)
    app.Get("/users", handlers.GetUsers)

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
