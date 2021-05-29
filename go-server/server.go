package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/listen", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Responding back",
		})
	})
	app.Post("/postdata", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Responding back",
		})
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen("0.0.0.0:5000"); err != nil {
		log.Panic(err)
	}

	select {
	case <-c:
		log.Println("Received second signal")
		break
	case <-time.After(time.Second * time.Duration(5)):
		break
	}
}
