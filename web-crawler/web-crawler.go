package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mut  sync.Mutex
	vals map[string]bool
}

func (c *Cache) contains(url string) bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.vals == nil {
		c.vals = make(map[string]bool)
	}
	if c.vals[url] {
		return true
	}
	c.vals[url] = true
	return false
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache, ret chan bool) {
	defer close(ret)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	children := make([]chan bool, len(urls))
	for i, u := range urls {
		children[i] = make(chan bool)
		if !cache.contains(u) {
			go Crawl(u, depth-1, fetcher, cache, children[i])
		} else {
			close(children[i])
		}
	}

	for i := range children {
		for s := range children[i] {
			ret <- s
		}
	}
	return
}

func main() {
	cache := new(Cache)
	ret := make(chan bool)
	Crawl("https://golang.org/", 4, fetcher, cache, ret)
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
