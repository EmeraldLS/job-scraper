package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// notworking
func ScrapeJobberman() {
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Println("Error making request: ", err)
		}
	})

	c.OnHTML("div[listing-cards-components]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)

	})

	c.Visit("https://www.jobberman.com/jobs")
}
