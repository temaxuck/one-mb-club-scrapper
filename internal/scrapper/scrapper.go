package scrapper

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func FetchBody(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch resource data: %v\n", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Resource responsed with non OK status: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Couldn't read body: %v\n", err)
	}
	
	return body, nil
}

func Scrap1MbClub() (urls []string, err error) {
	body, err := FetchBody(ONE_MB_CLUB_URL)

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

	urls = []string{}

	for _, url := range resourceLinkMatches {
		urls = append(urls, url[1])
	}
	return urls, nil
}
