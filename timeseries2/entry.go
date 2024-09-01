package timeseries2

import (
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
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

	Timeout        int32   `json:"timeout"`
	RateLimit      float64 `json:"rate-limit"`
	RateBurst      int32   `json:"rate-burst"`
	ControllerCode string  `json:"cc"`

	Route        string `json:"route"`
	RouteTo      string `json:"route-to"`
	RoutePercent int    `json:"route-percent"`
	RouteCode    string `json:"rc"`
}

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
		case RoutePercentName:
			e.RouteTo = values[i].(string)
		case RouteCodeName:
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

func (e Entry) Values() []any {
	return []any{
		e.StartTime,
		e.Duration,
		e.Traffic,
		e.CreatedTS,

		e.Region,
		e.Zone,
		e.SubZone,
		e.Host,
		e.InstanceId,

		e.RequestId,
		e.RelatesTo,
		e.Protocol,
		e.Method,
		e.From,
		e.To,
		e.Uri,
		e.Path,

		e.StatusCode,
		e.Encoding,
		e.Bytes,

		e.Timeout,
		e.RateLimit,
		e.RateBurst,
		e.ControllerCode,

		e.Route,
		e.RouteTo,
		e.RoutePercent,
		e.RouteCode,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func (e Entry) Origin() core.Origin {
	return core.Origin{
		Region:     e.Region,
		Zone:       e.Zone,
		SubZone:    e.SubZone,
		Host:       e.Host,
		InstanceId: "",
		Route:      e.Route,
	}
}
