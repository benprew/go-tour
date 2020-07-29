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

type UrlList struct {
	urls map[string]bool
	mux  sync.Mutex
}

var urlList UrlList

func notify(c chan int) {
	c <- 1
}

func seen(url string) bool {
	urlList.mux.Lock()
	defer urlList.mux.Unlock()

	if !urlList.urls[url] {
		urlList.urls[url] = true
		return false
	} else {
		return true
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, results chan int) {
	// Fetch URLs in parallel.
	//     Fetch URLs in parallel by using go routines and a channel to monitor
	//     when the sub-crawls are done.
	// Don't fetch the same URL twice.
	//     Uses a map to keep from fetching the same url more than once and uses a
	//     mutex to synchronize access to the map

	defer notify(results) // add a message to results to let caller know we're done.
	if depth <= 0 {
		return
	}
	if seen(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	c := make(chan int, len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c)
	}
	// wait for all sub-crawls to finish
	for i := 0; i < len(urls); i++ {
		fmt.Printf("read %d, len: %d url: %s\n", i, len(urls), url)
		<-c
	}
}

func main() {
	results := make(chan int, 1)
	urlList.urls = make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher, results)
	<-results // read 1 result from the first crawl
	fmt.Println("### Results ###")
	for n := range urlList.urls {
		fmt.Println(n)
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
