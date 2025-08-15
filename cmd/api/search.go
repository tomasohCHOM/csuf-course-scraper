package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://catalog.fullerton.edu"

type Course struct {
	Title         string
	Description   string
	Prerequisites []string
	Corequisites  []string
}

func GetCourse(url string) Course {
	c := colly.NewCollector()
	course := Course{}

	c.OnHTML("td.block_content", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("h1#course_preview_title").Text())
		course.Title = title
		courseContent := strings.Split(e.Text, "\n")
		fmt.Println(courseContent)
	})

	c.Visit(url)
	return course
}

func SearchCourses(courseQuery string) error {
	url := fmt.Sprintf("%s/search_advanced.php?cur_cat_oid=80&search_database=Search&search_db=Search&cpage=1&ecpage=1&ppage=1&spage=1&tpage=1&location=3&filter[keyword]=%s&filter[exact_match]=1", BASE_URL, courseQuery)
	c := colly.NewCollector()
	// Scrape all the courses based on the CSS selector
	c.OnHTML(`a[aria-expanded="false"]`, func(e *colly.HTMLElement) {
		courseUrl := fmt.Sprintf("%s/%s", BASE_URL, e.Attr("href"))
		GetCourse(courseUrl)
	})
	err := c.Visit(url)
	if err != nil {
		return fmt.Errorf("failed to visit URL: %w", err)
	}
	return nil
}
