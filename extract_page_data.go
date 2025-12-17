package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	h1 := getH1FromHTML(html)
	p := getFirstParagraphFromHTML(html)

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			H1:             h1,
			FirstParagraph: p,
		}
	}

	outgoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		outgoingLinks = nil
	}

	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		imageURLs = nil
	}

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: p,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
}
