package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("no website provided")
	}
	if len(args) > 1 {
		log.Fatal("too many arguments provided")
	}
	rawBaseURL := args[0]

	const maxConcurrency = 5
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		log.Fatalf("Error - configure: %v", err)
	}

	fmt.Printf("Starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()
}
