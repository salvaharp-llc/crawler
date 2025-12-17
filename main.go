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
	baseURL := args[0]

	fmt.Printf("Starting crawl of: %s\n", baseURL)

	pages := map[string]int{}
	crawlPage(baseURL, baseURL, pages)

	for key, val := range pages {
		fmt.Printf("- %s visited %d times\n", key, val)
	}
}
