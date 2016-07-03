package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fetchResult struct {
	url, body string
	urls      []string
	depth     int
	err       error
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	ch := make(chan *fetchResult)
	fetched := make(map[string]bool)
	fetch := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		ch <- &fetchResult{url, body, urls, depth, err}
	}

	go fetch(url, depth)
	fetched[url] = true

	for i := 1; i > 0; i-- {
		res := <-ch

		if res.err != nil {
			fmt.Println(res.err)
			continue
		}

		fmt.Printf("found: %s %q\n", res.url, res.body)

		if res.depth > 1 {
			for _, u := range res.urls {
				if !fetched[u] {
					go fetch(u, res.depth-1)
					fetched[u] = true
					i++
				}
			}
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
