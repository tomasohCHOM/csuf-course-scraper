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
	Prerequisites string
	Corequisites  string
}

func FilterText(text []string, courseQuery string) string {
	for _, s := range text {
		if strings.Contains(s, strings.Join(strings.Split(courseQuery, "+"), " ")) {
			return strings.TrimSpace(s)
		}
	}
	return ""
}

func GetRequisites(description string) (string, string) {
	var prerequisites string
	var corequisites string

	if strings.Contains(description, "Prerequisites") {
		prerequisites = strings.Split(description, "Prerequisites:")[1]
		prerequisites = strings.Split(prerequisites, ".")[0]
	}

	if strings.Contains(description, "Corerequisites") {
		corequisites = strings.Split(description, "Corerequisites:")[1]
		corequisites = strings.Split(corequisites, ".")[0]
	}

	return prerequisites, corequisites
}

func GetCourse(url string, courseQuery string) Course {
	c := colly.NewCollector()
	course := Course{}

	c.OnHTML("td.block_content", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("h1#course_preview_title").Text())

		text := strings.Split(e.Text, "\n")
		description := strings.Split(FilterText(text, courseQuery), title)[1]

		course.Title = title

		prerequisites, corequisites := GetRequisites(description)

		if len(prerequisites) > 0 {
			description = strings.Split(description, "Prerequisites:")[0]
		}
		if len(corequisites) > 0 {
			description = strings.Split(description, "Corequisites:")[0]
		}

		course.Description = strings.TrimSpace(description)
		course.Prerequisites = strings.TrimSpace(prerequisites)
		course.Corequisites = strings.TrimSpace(corequisites)
	})

	c.Visit(url)
	return course
}

func SearchCourse(courseQuery string) (Course, error) {
	url := fmt.Sprintf("%s/search_advanced.php?cur_cat_oid=80&search_database=Search&search_db=Search&cpage=1&ecpage=1&ppage=1&spage=1&tpage=1&location=3&filter[keyword]=%s&filter[exact_match]=1", BASE_URL, courseQuery)
	c := colly.NewCollector()
	course := Course{}

	c.OnHTML(`a[aria-expanded="false"]`, func(e *colly.HTMLElement) {
		courseUrl := fmt.Sprintf("%s/%s", BASE_URL, e.Attr("href"))
		course = GetCourse(courseUrl, courseQuery)
	})

	if err := c.Visit(url); err != nil {
		return Course{}, fmt.Errorf("failed to visit URL: %w", err)
	}
	return course, nil
}
