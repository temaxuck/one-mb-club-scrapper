package metrics

import (
	"time"
)

type Metrics struct {
	URL     string
	SizeKB  float
	FetchTS time.Time
	Status  string
}
