package main

import (
	"net/url"
	"strings"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
	ReqStatus      string
	Error          string
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
			ReqStatus:      "Success",
			Error:          err.Error(),
		}
	}

	var errors []string
	outgoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		outgoingLinks = nil
		errors = append(errors, err.Error())
	}

	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		imageURLs = nil
		errors = append(errors, err.Error())
	}

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: p,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
		ReqStatus:      "Success",
		Error:          strings.Join(errors, ";"),
	}
}
