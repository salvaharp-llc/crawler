package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if !cfg.isInternalLink(currentURL) {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizeURL: %v", err)
		return
	}

	if isFirst := cfg.addPageVisit(normalizedURL); !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		cfg.setFailedPage(normalizedURL, rawCurrentURL, err)
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	pageData := extractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	for _, u := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}
