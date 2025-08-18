package main

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

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

type Requisites struct {
	Prerequisites string
	Corequisites  string
}

func getRequisites(description string) Requisites {
	prereqRegex := regexp.MustCompile(`(?i)Prerequisites?:\s*([^.]*)`)
	coreqRegex := regexp.MustCompile(`(?i)Corequisites?:\s*([^.]*)`)

	r := Requisites{}

	if match := prereqRegex.FindStringSubmatch(description); match != nil {
		r.Prerequisites = strings.TrimSpace(match[1])
	}

	if match := coreqRegex.FindStringSubmatch(description); match != nil {
		r.Corequisites = strings.TrimSpace(match[1])
	}

	return r
}

var requisitesRegex = regexp.MustCompile(`(?i)(Pre|Co)requisites?:\s*[^.]*\.`)

func GetCourse(url string, courseQuery string) Course {
	c := colly.NewCollector()
	course := Course{}

	c.OnHTML("td.block_content", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.DOM.Find("h1#course_preview_title").Text())

		fullText := FilterText(strings.Split(e.Text, "\n"), courseQuery)
		descriptionParts := strings.Split(fullText, title)
		var description string
		if len(descriptionParts) > 1 {
			description = strings.TrimSpace(descriptionParts[1])
		}

		reqs := getRequisites(description)
		cleanDesc := strings.TrimSpace(requisitesRegex.ReplaceAllString(description, ""))

		course.Title = title
		course.Description = cleanDesc
		course.Prerequisites = reqs.Prerequisites
		course.Corequisites = reqs.Corequisites
	})

	c.Visit(url)
	return course
}
