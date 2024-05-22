package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Course struct {
	Title          string
	Description    string
	Prerequisites  []string
	Corerequisites []string
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("frontend/*.html")),
	}
}

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.Use(middleware.Logger())
	e.Renderer = createTemplate()

	var courses []Course
	c := colly.NewCollector()

	// Scrape all the courses based on the CSS selector
	c.OnHTML("li.acalog-course", func(e *colly.HTMLElement) {
		course := Course{}
		course.Title = e.ChildText("span")
		course.Description = "Course description"
		course.Prerequisites = []string{"CPSC 131"}
		course.Corerequisites = []string{course.Title}

		courses = append(courses, course)
	})

	data := struct{ Courses []Course }{Courses: courses}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", data)
	})

	// Visit the CSUF CPSC Catalog website
	c.Visit("https://catalog.fullerton.edu/preview_program.php?catoid=80&poid=38156&returnto=11049")
	e.Logger.Fatal(e.Start(":3000"))
}
