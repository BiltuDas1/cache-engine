package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/redirect"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func configureApp(app *fiber.App) {
	// Redirect from index.html to / (root)
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"index.html": "/",
		},
		StatusCode: 308,
	}))

	// Serve Static Content
	if runningInDocker {
		app.Get("/", static.New("/app/index.html"))
		app.Get("/*", static.New("/app"))
	} else {
		app.Get("/", static.New("./engine-ui/index.html"))
		app.Get("/*", static.New("./engine-ui"))
	}
}
