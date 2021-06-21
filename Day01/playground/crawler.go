package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeMap is safe to use concurrently.
type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

// checkVisited(url) checks if an url was visited already. If no, store it.
func (c *SafeMap) checkVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.v[url]; !ok {
		c.v[url] = true
		fmt.Println("New: ", url)
		return false
	}
	fmt.Println("Old: ", url)
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(wg *sync.WaitGroup, c *SafeMap, url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	if !c.checkVisited(url) {
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go Crawl(wg, c, u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	cache := &SafeMap{}
	cache.v = make(map[string]bool)
	Crawl(wg, cache, "https://golang.org/", 4, fetcher)
	wg.Wait()
	for k, v := range cache.v {
		fmt.Println("k=", k, " v=", v)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
