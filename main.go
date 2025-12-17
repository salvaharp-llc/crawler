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

	fmt.Printf("starting crawl of: %s\n", baseURL)

	htmlBody, err := getHTML(baseURL)
	if err != nil {
		log.Fatalf("couldn't get HTML: %v", err)
	}

	fmt.Println(htmlBody)

}
