package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://catalog.fullerton.edu"

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
