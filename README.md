# crawler

A small Go-based web crawler that visits pages from a base URL and writes a CSV report (`report.csv`) with extracted page data.

## Prerequisites ✅

- **Go 1.25** or later installed (the project uses `go 1.25.0` in `go.mod`)
- **Git**

## Clone the repository 🔧

```bash
git clone https://github.com/salvaharp-llc/crawler.git
cd crawler
```

## Build & Run ▶️

Run directly with `go run`:

```bash
go run . <URL> <maxConcurrency> <maxPages>
# example
go run . https://example.com 5 100
```

Or build a binary:

```bash
go build -o crawler .
./crawler https://example.com 5 100
```

The crawler writes results to `report.csv` in the repository root.

## Tests ✅

Run unit tests:

```bash
go test ./...
```

## Notes & Tips 💡

- Dependencies are managed by Go modules; running `go build` or `go test` will fetch them automatically.
- Adjust `maxConcurrency` and `maxPages` to control crawl parallelism and breadth.
- If you have issues or want to contribute, please open an issue or a pull request on GitHub.
