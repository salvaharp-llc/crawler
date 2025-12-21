package main

import (
	"net/url"
	"testing"
)

func TestGetRedirectFromHTML(t *testing.T) {
	cases := []struct {
		name    string
		html    string
		base    string
		wantURL string
		wantOk  bool
	}{
		{
			name:    "no meta",
			html:    `<html><head></head><body></body></html>`,
			base:    "https://example.com",
			wantURL: "",
			wantOk:  false,
		},
		{
			name:    "relative redirect",
			html:    `<html><head><meta http-equiv="refresh" content="0; url=/moved"></head><body></body></html>`,
			base:    "https://example.com",
			wantURL: "https://example.com/moved",
			wantOk:  true,
		},
		{
			name:    "absolute redirect with quotes",
			html:    `<html><head><meta http-equiv="refresh" content="2;url='https://other.com/'"></head><body></body></html>`,
			base:    "https://example.com",
			wantURL: "https://other.com/",
			wantOk:  true,
		},
		{
			name:    "malformed content no url",
			html:    `<html><head><meta http-equiv="refresh" content="5"></head></html>`,
			base:    "https://example.com",
			wantURL: "",
			wantOk:  false,
		},
		{
			name:    "empty url",
			html:    `<html><head><meta http-equiv="refresh" content="0; url=' '"></head></html>`,
			base:    "https://example.com",
			wantURL: "",
			wantOk:  false,
		},
		{
			name:    "invalid url",
			html:    `<html><head><meta http-equiv="refresh" content="0; url=://invalid"></head></html>`,
			base:    "https://example.com",
			wantURL: "",
			wantOk:  false,
		},
		{
			name:    "malformed content no semicolon",
			html:    `<html><head><meta http-equiv="refresh" content="0 url=/moved"></head></html>`,
			base:    "https://example.com",
			wantURL: "https://example.com/moved",
			wantOk:  true,
		},
		{
			name:    "case variation url",
			html:    `<html><head><meta http-equiv="refresh" content="0; uRL=/moved"></head><body></body></html>`,
			base:    "https://example.com",
			wantURL: "https://example.com/moved",
			wantOk:  true,
		},
		{
			name:    "whitespace variation",
			html:    `<html><head><meta http-equiv="refresh" content="0; url = /moved"></head><body></body></html>`,
			base:    "https://example.com",
			wantURL: "https://example.com/moved",
			wantOk:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			base, err := url.Parse(tc.base)
			if err != nil {
				t.Fatalf("invalid base URL in test: %v", err)
			}
			got, ok := getRedirectFromHTML(tc.html, base)
			if got != tc.wantURL {
				t.Errorf("url: want %q, got %q", tc.wantURL, got)
			}
			if ok != tc.wantOk {
				t.Errorf("ok: want %v, got %v", tc.wantOk, ok)
			}
		})
	}
}
