package main

import (
	"log/slog"
	"os"

	"github.com/EwanGreer/blog/internal/templator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(e echo.Context) error {
		dataFields := map[string]any{
			"title": "My Blog",
		}

		return templator.New().Page("index").Execute(e.Response(), dataFields)
	})

	e.GET("/index/blogs/list", func(e echo.Context) error {
		// TODO: call db to get a paginated list of blogs.

		dataFields := map[string]any{
			"Blogs": []map[string]string{
				{
					"BlogTitle":       "How This Website Was built!",
					"CardDescription": "A brief explaination of how this website was built.",
					"Slug":            "templ-slug",
				},
				{
					"BlogTitle":       "What I'm Building Next!",
					"CardDescription": "What I'm working on now, and in the future.",
					"Slug":            "templ-slug",
				},
				{
					"BlogTitle":       "What I'm Building Next!",
					"CardDescription": "What I'm working on now, and in the future.",
					"Slug":            "templ-slug",
				},
				{
					"BlogTitle":       "What I'm Building Next!",
					"CardDescription": "What I'm working on now, and in the future.",
					"Slug":            "templ-slug",
				},
				{
					"BlogTitle":       "What I'm Building Next!",
					"CardDescription": "What I'm working on now, and in the future.",
					"Slug":            "templ-slug",
				},
				{
					"BlogTitle":       "What I'm Building Next!",
					"CardDescription": "What I'm working on now, and in the future.",
					"Slug":            "templ-slug",
				},
			},
		}

		return templator.New().Partial("blogs").Execute(e.Response(), dataFields)
	})

	if err := e.Start(":3000"); err != nil {
		slog.Error("Service shutting down", "err", err)
		os.Exit(1)
	}
}
