package metrics

import (
	"fmt"
	"time"

	"github.com/temaxuck/one-mb-club-scrapper/internal/scrapper"
)

type Metrics struct {
	URL     string
	SizeKB  float32
	FetchDuration time.Duration
	Status  string
}

func GatherMetrics(url string) Metrics {
	start_time := time.Now()
	body, err := scrapper.FetchBody(url)
	end_time := time.Now()
	if err != nil {
		return Metrics{
			URL:    url,
			FetchDuration: end_time.Sub(start_time),
			Status: fmt.Sprintf("ERROR: %v", err),
		}
	}

	return Metrics{
		URL:     url,
		SizeKB:  float32(len(body)) / 1024,
		FetchDuration: end_time.Sub(start_time),
		Status:  "OK",
	}
}

