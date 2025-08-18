package main

import (
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

func renderCourseError(c echo.Context, msg string) error {
	return c.Render(http.StatusOK, "course.html", map[string]any{
		"Error": msg,
	})
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
			return renderCourseError(c, "Query must be non-empty")
		}

		parts := strings.Fields(query)
		if len(parts) > 2 {
			return renderCourseError(c, "Invalid query string")
		}

		parsedQuery := strings.ToUpper(strings.Join(parts, "+"))
		course, err := SearchCourse(parsedQuery)
		if err != nil {
			return renderCourseError(c, err.Error())
		}

		return c.Render(http.StatusOK, "course.html", map[string]any{
			"Course": course,
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}
