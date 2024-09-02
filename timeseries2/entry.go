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
		case startTimeName:
			e.StartTime = values[i].(time.Time)
		case durationName:
			e.Duration = values[i].(int64)
		case trafficName:
			e.Traffic = values[i].(string)
		case createdTSName:
			e.CreatedTS = values[i].(time.Time)

		case regionName:
			e.Region = values[i].(string)
		case zoneName:
			e.Zone = values[i].(string)
		case subZoneName:
			e.SubZone = values[i].(string)
		case hostName:
			e.Host = values[i].(string)
		case instanceIdName:
			e.InstanceId = values[i].(string)

		case requestIdName:
			e.RequestId = values[i].(string)
		case relatesToName:
			e.RelatesTo = values[i].(string)
		case protocolName:
			e.Protocol = values[i].(string)
		case methodName:
			e.Method = values[i].(string)
		case fromName:
			e.From = values[i].(string)
		case toName:
			e.To = values[i].(string)
		case uriName:
			e.Uri = values[i].(string)
		case pathName:
			e.Path = values[i].(string)

		case statusCodeName:
			e.StatusCode = values[i].(int32)
		case encodingName:
			e.Encoding = values[i].(string)
		case bytesName:
			e.Bytes = values[i].(int64)

		case routeName:
			e.Route = values[i].(string)
		case routeToName:
			e.RouteTo = values[i].(string)
		case routePercentName:
			e.RouteTo = values[i].(string)
		case routeCodeName:
			e.RouteTo = values[i].(string)

		case timeoutName:
			e.Timeout = values[i].(int32)
		case rateLimitName:
			e.RateLimit = values[i].(float64)
		case rateBurstName:
			e.RateBurst = values[i].(int32)
		case controllerCodeName:
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
