package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("frontend/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.Renderer = createTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/course", func(c echo.Context) error {
		query := c.QueryParam("q")
		if query == "" {
			return c.Render(http.StatusBadRequest, "course.html", map[string]any{
				"Error": "Query must be non-empty",
			})
		}

		queryContents := strings.Split(query, " ")
		if len(queryContents) > 2 {
			return c.Render(http.StatusBadRequest, "course.html", map[string]any{
				"Error": "Invalid query string.",
			})
		}
		parsedQuery := strings.Join(queryContents, "+")
		fmt.Println(parsedQuery)

		course, err := SearchCourse(parsedQuery)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "course.html", map[string]any{
				"Error": "Error fetching course data. Try again later.",
			})
		}
		return c.Render(http.StatusOK, "course.html", map[string]any{
			"Course": course,
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
