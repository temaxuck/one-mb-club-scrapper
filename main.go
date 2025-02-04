package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

// NOTE: Well, I kinda have an idea of scrapping the websites from 1 MB Club (https://1mb.club/) and making something out of it.
// TODO: Elaborate the idea and probably this could be some funny project.

const ONE_MB_CLUB_URL = "https://1mb.club/"
const MAX_RESOURCE_COUNT = 60

func fetchBody(url string) ([]byte, error) {
	r, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch resource data: %v\n", err)
	}
	defer r.Body.Close()

	fmt.Printf("Status of %s: %s\n", url, r.Status);
	
	
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Resource responsed with non OK status: %v\n", r.StatusCode)
	}

	return io.ReadAll(r.Body)
}

func fetchResourceURLs() (urls []string, err error) {
	body, err := fetchBody(ONE_MB_CLUB_URL)

	if err != nil {
		return nil, err
	}

	tableContentsPattern := regexp.MustCompile(`<tbody\s*id="container">([\w\W\s]*?)</tbody>`)
	tableContentsMatch := tableContentsPattern.FindStringSubmatch(string(body))
	if len(tableContentsMatch) < 2 {
		return nil, fmt.Errorf("Couldn't parse %s: No contents of `<tbody id=\"container\"></tbody>` was found.", ONE_MB_CLUB_URL)
	}
	tableContents := tableContentsMatch[1]

	resourceLinkPattern := regexp.MustCompile(`<a[^>]*href="([^"]+)"[^>]*>`)
	resourceLinkMatches := resourceLinkPattern.FindAllStringSubmatch(tableContents, -1)

	result := []string{}

	for _, url := range resourceLinkMatches {
		result = append(result, url[1])
	}
	return result, nil
}

func sync_fetch(urls []string) {
	for i := 0; i < len(urls) && i < MAX_RESOURCE_COUNT; i++ {
		url := urls[i]
		body, err := fetchBody(url)
		if err != nil {
			fmt.Printf("%d: WARNING: Couldn't read body from '%s'. Cause: %v\n", i, url, err)
			continue
		}
		fmt.Printf("%d: Fetched data from '%s'\n", i, url, string(body))
	}
}

func async_fetch(urls []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(urls) && i < MAX_RESOURCE_COUNT; i++ {
		url := urls[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := fetchBody(url)
			if err != nil {
				fmt.Printf("%d: WARNING: Couldn't read body from '%s'. Cause: %v\n", i, url, err)
				return
			}
			fmt.Printf("%d: Fetched data from '%s'\n", i, url)
		}()
	}

	wg.Wait()
}

func main() {
	// Measure execution time
	start_time := time.Now()

	urls, err := fetchResourceURLs()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	async_fetch(urls)

	fmt.Printf("Execution time: %.2f seconds\n", time.Now().Sub(start_time).Seconds())
}
