package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		log.Fatal("usage: go run . <URL> <maxConcurrency> <maxPages>")
	}
	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("invalid maxConcurrency value: %v", err)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("invalid maxPages value: %v", err)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		log.Fatalf("Error - configure: %v", err)
	}

	fmt.Printf("Starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	if err := writeCSVReport(cfg.pages, "report.csv"); err != nil {
		fmt.Printf("Error - writeCSVReport: %v\n", err)
	}
}
