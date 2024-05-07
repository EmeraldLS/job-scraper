package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var jobMagScraper = JobMagScraper{}

func TestScrapeByLocation(t *testing.T) {
	jobs := jobMagScraper.ScrapeByLocation("lagos", 1)
	b, err := json.MarshalIndent(jobs, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

func TestScrapeJobs(t *testing.T) {
	jobs := jobMagScraper.ScrapeJobs(JobMagFilter{}, 1)
	b, err := json.MarshalIndent(jobs, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
