package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getRedirectFromHTML(html string, baseURL *url.URL) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", false
	}

	selection := doc.Find("meta[http-equiv='refresh']")
	if selection.Length() == 0 {
		return "", false
	}

	var redirect string
	selection.EachWithBreak(func(_ int, s *goquery.Selection) bool {
		content, exists := s.Attr("content")
		if !exists {
			return true
		}

		if strings.Contains(strings.ToLower(content), "url") {
			redirect = content
			return false // break
		}
		return true
	})
	if redirect == "" {
		return "", false
	}

	rawRedirectURL := strings.Split(redirect, "=")[1]
	rawRedirectURL = strings.TrimFunc(rawRedirectURL, func(r rune) bool {
		return r == '\'' || r == '"' || r == ' '
	})
	if rawRedirectURL == "" {
		return "", false
	}

	redirectURL, err := url.Parse(rawRedirectURL)
	if err != nil {
		return "", false
	}

	absolute := baseURL.ResolveReference(redirectURL)
	return absolute.String(), true
}
