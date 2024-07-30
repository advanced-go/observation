package access1

import (
	"time"
)

// Entry - timeseries access log struct
type Entry struct {
	StartTime time.Time `json:"start-time"`
	Duration  int64     `json:"duration"`
	Traffic   string    `json:"traffic"`
	CreatedTS time.Time `json:"created-ts"`

	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`

	RequestId string `json:"request-id"`
	RelatesTo string `json:"relates-to"`
	Protocol  string `json:"proto"`
	Method    string `json:"method"`
	From      string `json:"from"`
	To        string `json:"to"`
	Url       string `json:"url"`
	Path      string `json:"path"`

	StatusCode int32  `json:"status-code"`
	Encoding   string `json:"encoding"`
	Bytes      int64  `json:"bytes"`

	Route      string  `json:"route"`
	RouteTo    string  `json:"route-to"`
	Timeout    int32   `json:"timeout"`
	RateLimit  float64 `json:"rate-limit"`
	RateBurst  int32   `json:"rate-burst"`
	ReasonCode string  `json:"rc"`
}
