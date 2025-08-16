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
			return c.String(http.StatusBadRequest, "Missing query")
		}

		queryContents := strings.Split(query, " ")
		if len(queryContents) > 2 {
			return c.String(http.StatusBadRequest, "Invalid query string")
		}
		parsedQuery := strings.Join(queryContents, "+")
		fmt.Println(parsedQuery)

		course, err := SearchCourses(parsedQuery)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch course")
		}
		return c.Render(http.StatusOK, "course.html", course)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
