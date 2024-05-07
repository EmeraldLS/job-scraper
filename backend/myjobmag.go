package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type JobMagScraper struct{}

type JobMag struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	URL         string `json:"url"`
}

type JobMagFilter struct {
	Field      string
	Industry   string
	Location   string
	Experience string
	Education  string
}

var c = colly.NewCollector()

func (*JobMagScraper) ScrapeByLocation(state string, page int) []JobMag {
	var jobs []JobMag

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

	c.OnHTML(".mag-b .cat-h1", func(h *colly.HTMLElement) {
		text := h.Text
		fmt.Println("Text = ", text)
	})

	c.OnHTML(".job-list", func(h *colly.HTMLElement) {
		h.ForEach(".job-list-li", func(_ int, j *colly.HTMLElement) {
			j.ForEach(".job-info ul", func(i int, k *colly.HTMLElement) {
				url := k.ChildAttr(".mag-b a", "href")
				title := k.ChildText(".mag-b h2")
				desc := k.ChildText(".job-desc")
				date := k.ChildText("#job-date")

				if title != "" && desc != "" && date != "" {
					job := JobMag{
						Title:       title,
						Description: desc,
						Date:        date,
						URL:         "https://www.myjobmag.com" + url,
					}

					jobs = append(jobs, job)
				}
			})
		})
	})

	c.Visit(fmt.Sprintf("https://www.myjobmag.com/jobs-location/%s/%d", state, page))

	return jobs
}

func (*JobMagScraper) ScrapeJobs(filter JobMagFilter, page int) []JobMag {
	var jobs []JobMag

	c := colly.NewCollector()
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

	c.OnHTML(".mag-b h1", func(h *colly.HTMLElement) {
		pageTitle := h.Text
		fmt.Println(pageTitle)
	})

	c.OnHTML(".job-list", func(h *colly.HTMLElement) {
		h.ForEach(".job-list-li", func(_ int, j *colly.HTMLElement) {
			j.ForEach(".job-info ul", func(i int, k *colly.HTMLElement) {
				url := k.ChildAttr(".mag-b a", "href")
				title := k.ChildText(".mag-b h2")
				desc := k.ChildText(".job-desc")
				date := k.ChildText("#job-date")

				if title != "" && desc != "" && date != "" {
					job := JobMag{
						Title:       title,
						Description: desc,
						Date:        date,
						URL:         "https://www.myjobmag.com" + url,
					}

					jobs = append(jobs, job)
				}
			})
		})
	})

	c.Visit(fmt.Sprintf("https://www.myjobmag.com/search/jobs?currpage=%d&field=%s&industry=%s&location=%sexperience=%s&education=%s", page, filter.Field, filter.Industry, filter.Location, filter.Experience, filter.Education))

	return jobs
}
