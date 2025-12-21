package main

import (
	"net/url"
)

type ReqStatus string

const (
	ReqStatusSuccess  ReqStatus = "Success"
	ReqStatusFailed   ReqStatus = "Failed"
	ReqStatusRedirect ReqStatus = "Redirect"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
	ReqStatus      ReqStatus
	Error          string
}

func extractPageData(html, pageURL string) PageData {
	reqStatus := ReqStatusSuccess

	h1 := getH1FromHTML(html)
	p := getFirstParagraphFromHTML(html)

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			H1:             h1,
			FirstParagraph: p,
			ReqStatus:      reqStatus,
			Error:          err.Error(),
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

	redirectURL, isRedirect := getRedirectFromHTML(html, baseURL)
	if isRedirect {
		reqStatus = ReqStatusRedirect
		outgoingLinks = append(outgoingLinks, redirectURL)
	}

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: p,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
		ReqStatus:      reqStatus,
	}
}
