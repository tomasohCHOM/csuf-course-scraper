package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

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
		err := SearchCourses("CPSC+240")
		if err != nil {
			log.Fatal(err)
		}
		return c.Render(http.StatusOK, "index.html", []int{})
	})
	e.Logger.Fatal(e.Start(":3000"))
}
