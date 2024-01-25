package main

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// Options of the tool
type Options struct {
	ListenAddress string
	Folder        string
}

func main() {
	options := ParseOptions()
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Static("/", options.Folder)

	app.Listen(options.ListenAddress)
}

func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.ListenAddress, "listen", "0.0.0.0:8000", "Address:Port")
	currentPath := "."
	if p, err := os.Getwd(); err == nil {
		currentPath = p
	}
	flag.StringVar(&options.Folder, "path", currentPath, "Folder")
	flag.Parse()

	return options
}
