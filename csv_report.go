package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to CSV")
		return nil
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("couldn't create file: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
		"request_status",
		"error",
	}
	if err := w.Write(header); err != nil {
		return fmt.Errorf("couldn't write header: %v", err)
	}

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, normalizedURL := range keys {
		p := pages[normalizedURL]
		outgoing := strings.Join(p.OutgoingLinks, ";")
		images := strings.Join(p.ImageURLs, ";")
		row := []string{
			p.URL,
			p.H1,
			p.FirstParagraph,
			outgoing,
			images,
			string(p.ReqStatus),
			p.Error,
		}
		if err := w.Write(row); err != nil {
			return fmt.Errorf("write row for %s: %w", p.URL, err)
		}
	}

	fmt.Printf("Report written to %s\n", filename)
	return nil
}
