package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name: "one h1",
			html: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "two h1",
			html: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
	<h1>Farewell</h1>
  </body>
</html>
			`,
			expected: "Welcome to Boot.dev",
		},
		{
			name: "no h1",
			html: `
<html>
  <body>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name: "paragraphs ouside main",
			html: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
	<p>Learn to code by building real projects.</p>
    <p>This is the second paragraph.</p>
    <main>
    </main>
  </body>
</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "paragraphs inside main",
			html: `
<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
	<h1>Farewell</h1>
  </body>
</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "paragraphs inside and outside main",
			html: `
<html>
  <body>
  	<p>Not this one.</p>
	<p>Not this one either.</p>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "no paragraphs",
			html: `
<html>
  <body>
    <main>
    </main>
  </body>
</html>
			`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
