package timeseries1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/observation/common"
	"github.com/advanced-go/stdlib/access"
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
	Location  string `json:"location"`
	Protocol  string `json:"proto"`
	Method    string `json:"method"`
	From      string `json:"from"`
	To        string `json:"to"`
	Uri       string `json:"uri"`
	Path      string `json:"path"`
	Query     string `json:"query"`

	StatusCode int32  `json:"status-code"`
	Encoding   string `json:"encoding"`
	Bytes      int64  `json:"bytes"`

	Route          string  `json:"route"`
	RouteTo        string  `json:"route-to"`
	RoutePercent   int     `json:"route-percent"`
	RouteCode      string  `json:"rc"`
	Timeout        int32   `json:"timeout"`
	RateLimit      float64 `json:"rate-limit"`
	RateBurst      int32   `json:"rate-burst"`
	ControllerCode string  `json:"cc"`
}

var (
	safeEntry = common.NewSafe()
	entryData = []Entry{
		{Region: "us-west1", Zone: "a", Host: "www.host1.com", Duration: 100, Traffic: access.IngressTraffic, Route: "host", Timeout: 2000, RateLimit: 98.5, RateBurst: 10, ControllerCode: "RL", StartTime: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-west1", Zone: "a", Host: "www.host2.com", Duration: 85, Traffic: access.IngressTraffic, Route: "host", Timeout: 1500, RateLimit: 100, RateBurst: 10, ControllerCode: "", StartTime: time.Date(2024, 6, 10, 7, 120, 55, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host1.com", Duration: 200, Traffic: access.IngressTraffic, Route: "host", Timeout: 300, RateLimit: 98.5, RateBurst: 10, ControllerCode: "RL", StartTime: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{Region: "us-central1", Zone: "c", Host: "www.host2.com", Duration: 750, Traffic: access.IngressTraffic, Route: "host", Timeout: 500, RateLimit: 100, RateBurst: 10, ControllerCode: "TO", StartTime: time.Date(2024, 6, 10, 7, 120, 55, 0, time.UTC)},
	}
)

func (Entry) Scan(columnNames []string, values []any) (e Entry, err error) {
	for i, name := range columnNames {
		switch name {
		case StartTimeName:
			e.StartTime = values[i].(time.Time)
		case DurationName:
			e.Duration = values[i].(int64)
		case TrafficName:
			e.Traffic = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case RegionName:
			e.Region = values[i].(string)
		case ZoneName:
			e.Zone = values[i].(string)
		case SubZoneName:
			e.SubZone = values[i].(string)
		case HostName:
			e.Host = values[i].(string)
		case InstanceIdName:
			e.InstanceId = values[i].(string)

		case RequestIdName:
			e.RequestId = values[i].(string)
		case RelatesToName:
			e.RelatesTo = values[i].(string)
		case ProtocolName:
			e.Protocol = values[i].(string)
		case MethodName:
			e.Method = values[i].(string)
		case FromName:
			e.From = values[i].(string)
		case ToName:
			e.To = values[i].(string)
		case UriName:
			e.Uri = values[i].(string)
		case PathName:
			e.Path = values[i].(string)

		case StatusCodeName:
			e.StatusCode = values[i].(int32)
		case EncodingName:
			e.Encoding = values[i].(string)
		case BytesName:
			e.Bytes = values[i].(int64)

		case RouteName:
			e.Route = values[i].(string)
		case RouteToName:
			e.RouteTo = values[i].(string)

		case TimeoutName:
			e.Timeout = values[i].(int32)
		case RateLimitName:
			e.RateLimit = values[i].(float64)
		case RateBurstName:
			e.RateBurst = values[i].(int32)
		case ControllerCodeName:
			e.ControllerCode = values[i].(string)
		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (a Entry) Values() []any {
	return []any{
		a.StartTime,
		a.Duration,
		a.Traffic,
		a.CreatedTS,

		a.Region,
		a.Zone,
		a.SubZone,
		a.Host,
		a.InstanceId,

		a.RequestId,
		a.RelatesTo,
		a.Protocol,
		a.Method,
		a.From,
		a.To,
		a.Uri,
		a.Path,

		a.StatusCode,
		a.Encoding,
		a.Bytes,

		a.Route,
		a.RouteTo,
		a.Timeout,
		a.RateLimit,
		a.RateBurst,
		a.ControllerCode,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
