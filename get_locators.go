package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(html string, baseURL *url.URL) ([]string, error) {
	return getLocatorsFromHTML(html, "a", "href", baseURL)
}

func getImagesFromHTML(html string, baseURL *url.URL) ([]string, error) {
	return getLocatorsFromHTML(html, "img", "src", baseURL)
}

func getLocatorsFromHTML(html, element, attribute string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var urls []string
	selector := fmt.Sprintf("%s[%s]", element, attribute)
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		att, _ := s.Attr(attribute)
		if att = strings.TrimSpace(att); att == "" {
			return
		}

		u, err := url.Parse(att)
		if err != nil {
			fmt.Printf("couldn't parse %s: %v\n", attribute, err)
			return
		}

		absolute := baseURL.ResolveReference(u)
		urls = append(urls, absolute.String())
	})
	return urls, nil
}
