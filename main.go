package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"golang.org/x/sync/errgroup"
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

	registerRoutes(app, options)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return app.Listen(options.ListenAddress)
	})

	// Wait for shutdown signal
	g.Go(func() error {
		<-ctx.Done()
		return app.Shutdown()
	})

	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
	}

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

func registerRoutes(app *fiber.App, options *Options) {
	// Root route
	app.Static("/", options.Folder, fiber.Static{Browse: true, Download: true})
}
