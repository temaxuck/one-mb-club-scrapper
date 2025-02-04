package scrapper

import (
	"io"
	"net/http"
)

func FetchBody(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Couldn't fetch resource data: %v\n", err)
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Resource responsed with non OK status: %d\n", r.StatusCode)
	}

	return io.ReadAll(r.Body), nil
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

	result := []string{}

	for _, url := range resourceLinkMatches {
		result = append(result, url[1])
	}
	return result, nil
}
