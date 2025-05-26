package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const BASE_URL = "https://catalog.fullerton.edu"

func SearchCourses(courseQuery string) error {
	url := fmt.Sprintf("%s/search_advanced.php?cur_cat_oid=80&search_database=Search&search_db=Search&cpage=1&ecpage=1&ppage=1&spage=1&tpage=1&location=3&filter[keyword]=%s&filter[exact_match]=1", BASE_URL, courseQuery)
	fmt.Println(url)
	c := colly.NewCollector()
	// Scrape all the courses based on the CSS selector
	c.OnHTML(`a[aria-expanded="false"]`, func(e *colly.HTMLElement) {
		fmt.Println(e)
	})
	err := c.Visit(url)
	if err != nil {
		return fmt.Errorf("Failed to visit URL: %w", err)
	}
	return nil
}
