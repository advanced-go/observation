package access1

import (
	"errors"
	"fmt"
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
		case UrlName:
			e.Url = values[i].(string)
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
		case ReasonCodeName:
			e.ReasonCode = values[i].(string)
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
		a.Url,
		a.Path,

		a.StatusCode,
		a.Encoding,
		a.Bytes,

		a.Route,
		a.RouteTo,
		a.Timeout,
		a.RateLimit,
		a.RateBurst,
		a.ReasonCode,
	}
}

func (Entry) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}
