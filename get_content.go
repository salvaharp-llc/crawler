package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	h1 := doc.Find("h1").First().Text()
	return strings.TrimSpace(h1)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	pSel := doc.Find("main").Find("p")
	if pSel.Length() == 0 {
		pSel = doc.Find("p")
	}
	p := pSel.First().Text()
	return strings.TrimSpace(p)
}
