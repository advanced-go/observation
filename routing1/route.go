package routing1

import (
	"errors"
	"fmt"
	"github.com/advanced-go/observation/common"
	"github.com/advanced-go/stdlib/core"
	"net/url"
	"time"
)

// TODO :

const (
	RouteIdName      = "route_id"
	TrafficName      = "traffic"
	RouteName        = "name"
	AuthorityName    = "authority"
	TimeoutName      = "timeout"
	RateLimitingName = "rate_limiting"
)

var (
	safeRoute = common.NewSafe()
	routeData = []Route{
		{EntryId: 1, RouteId: 1, Traffic: "egress", RouteName: "google-search", Authority: "github/advanced-go/observation", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, RouteId: 2, Traffic: "egress", RouteName: "google-search", Authority: "github/advanced-go/observation", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Route - need to find some way to avoid duplicate processing, as host will start and stop due to
// Kubernetes scaling
// Updating is allowed as changes should be minimum
type Route struct {
	EntryId   int       `json:"entry-id"`
	RouteId   int       `json:"route-id"`
	CreatedTS time.Time `json:"created-ts"`
	UpdatedTS time.Time `json:"created-ts"`

	Traffic      string `json:"traffic"` // Ingress or egress
	RouteName    string `json:"route"`
	Authority    string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	Version      string `json:"version"`   // Only for ingress
	RateLimiting bool   `json:"rate-limiting"`

	// Re-evalute if needed. Host can either be static or a preference for dynamic routing
	//Static       string `json:"static"`
	Host string `json:"host"` // Enables static routing.

}

/*
func (e Route) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
}


*/

func (Route) Scan(columnNames []string, values []any) (e Route, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case RouteIdName:
			e.EntryId = values[i].(int)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)
		case UpdatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case TrafficName:
			e.Traffic = values[i].(string)
		case RouteName:
			e.RouteName = values[i].(string)
		case AuthorityName:
			e.Authority = values[i].(string)

		case TimeoutName:
			e.Timeout = values[i].(int)
		case RateLimitingName:
			e.RateLimiting = values[i].(bool)
		case HostName:
			e.Host = values[i].(string)

		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Route) Values() []any {
	return []any{
		e.EntryId,
		e.RouteId,
		e.UpdatedTS,
		e.CreatedTS,

		e.Traffic,
		e.RouteName,
		e.Authority,
		e.Timeout,
		e.RateLimiting,
		e.Host,
	}
}

func (Route) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validRoute(values url.Values, e Entry) bool {
	if values == nil {
		return false
	}
	filter := core.NewOrigin(values)
	target := core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
	if !core.OriginMatch(target, filter) {
		return false
	}
	// Additional filtering
	return true
}
