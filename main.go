package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Course struct {
	title          string
	description    string
	prerequisites  []string
	corerequisites []string
}

func main() {
	var courses []Course
	c := colly.NewCollector()

	// Scrape all the courses based on the CSS selector
	c.OnHTML("li.acalog-course", func(e *colly.HTMLElement) {
		course := Course{}
		course.title = e.ChildText("span")
		course.description = "Course description"
		course.prerequisites = []string{"CPSC 131"}
		course.corerequisites = []string{course.title}

		courses = append(courses, course)
	})

	// Print all the courses after scraping them
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(courses)
	})

	// Visit the CSUF CPSC Catalog website
	c.Visit("https://catalog.fullerton.edu/preview_program.php?catoid=80&poid=38156&returnto=11049")
}
