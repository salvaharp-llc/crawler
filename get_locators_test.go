package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	base, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("failed to parse base url in test setup: %v", err)
	}
	tests := []struct {
		name     string
		html     string
		base     *url.URL
		expected []string
	}{
		{
			name:     "no a",
			html:     `<html><body><p>Not a link</p></body></html>`,
			base:     base,
			expected: nil,
		},
		{
			name:     "invalid href",
			html:     `<html><body><a href=":\\invalidURL"></a></body></html>`,
			base:     base,
			expected: nil,
		},
		{
			name:     "single relative href",
			html:     `<html><body><a href="/about">About</a></body></html>`,
			base:     base,
			expected: []string{"http://example.com/about"},
		},
		{
			name:     "single absolute href",
			html:     `<html><body><a href="http://example.com/search?q=go">Google</a></body></html>`,
			base:     base,
			expected: []string{"http://example.com/search?q=go"},
		},
		{
			name: "multiple hrefs",
			html: `<html><body>
		<a href="/about">About</a>
		<a href="http://example.com/search?q=go">Google</a>
	</body></html>`,
			base: base,
			expected: []string{
				"http://example.com/about",
				"http://example.com/search?q=go",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.html, tc.base)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	base, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("failed to parse base url in test setup: %v", err)
	}
	tests := []struct {
		name     string
		html     string
		base     *url.URL
		expected []string
	}{
		{
			name:     "no images",
			html:     `<html><body><p>Not an image</p></body></html>`,
			base:     base,
			expected: nil,
		},
		{
			name:     "invalid href",
			html:     `<html><body><img src=":\\invalidURL" alt="logo"/></body></html>`,
			base:     base,
			expected: nil,
		},
		{
			name:     "single relative src",
			html:     `<html><body><img src="/img/logo.png" alt="logo"/></body></html>`,
			base:     base,
			expected: []string{"http://example.com/img/logo.png"},
		},
		{
			name:     "single absolute src",
			html:     `<html><body><img src="http://cdn.example.com/assets/hero.jpg" alt="hero"/></body></html>`,
			base:     base,
			expected: []string{"http://cdn.example.com/assets/hero.jpg"},
		},
		{
			name: "multiple srcs",
			html: `<html><body>
		<img src="/logo.png" alt="Logo">
		<img src="https://cdn.boot.dev/banner.jpg">
	</body></html>`,
			base: base,
			expected: []string{
				"http://example.com/logo.png",
				"https://cdn.boot.dev/banner.jpg",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getImagesFromHTML(tc.html, tc.base)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
