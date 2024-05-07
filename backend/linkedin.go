package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/storage"
)

// notworking
func ScrapeLinkedIn() {
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	stg := &storage.InMemoryStorage{}

	c.SetStorage(stg)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited: ", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Println("Error visiting linked: ", err)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code: ", r.StatusCode)
	})

	c.Visit("https://www.linkedin.com/pulse/web-scraping-task-scheduling-golang-wilson-wei-ming-tan")

}
