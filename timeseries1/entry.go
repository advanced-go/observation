package timeseries1

import "time"

// Entry - timeseries1 access log struct
type Entry struct {
	StartTime time.Time
	Duration  int64
	Traffic   string

	Region     string
	Zone       string
	SubZone    string
	Host       string
	InstanceId string

	RequestId string
	RelatesTo string
	Protocol  string
	Method    string
	Authority string
	Url       string
	Path      string

	StatusCode int32
	Encoding   string
	Bytes      int64

	Route          string
	RouteTo        string
	Threshold      int
	ThresholdFlags string
}

type EntryV2 struct {
	CustomerId     string
	StartTime      time.Time
	Duration       int64
	DurationString string
	Traffic        string

	Region     string
	Zone       string
	SubZone    string
	Service    string
	InstanceId string
	RouteName  string

	RequestId string
	Url       string
	Protocol  string
	Method    string
	Host      string
	Path      string

	StatusCode  int32
	BytesSent   int64
	StatusFlags string

	Timeout        int32
	RateLimit      float64
	RateBurst      int32
	Retry          bool
	RetryRateLimit float64
	RetryRateBurst int32
	Failover       bool
	Proxy          bool
}
