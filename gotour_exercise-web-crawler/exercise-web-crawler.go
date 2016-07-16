package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fetchResult struct {
	url, body string
	urls      []string
	err       error
	depth     int
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	fetched := map[string]bool{url: true}
	urls := []string{url}
	ch := make(chan []string)
	fnCrawl := func(url string) {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			ch <- []string{}
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		ch <- urls
	}

	for i := 0; i < depth; i++ {
		var next []string

		for _, u := range urls {
			go fnCrawl(u)
		}

		for j := 0; j < len(urls); j++ {
			res := <-ch
			for _, r := range res {
				if !fetched[r] {
					fetched[r] = true
					next = append(next, r)
				}
			}
		}

		urls = next
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

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
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
