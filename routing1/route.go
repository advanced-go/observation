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
	RouteIdName   = "route_id"
	AuthorityName = "authority"
	VersionName   = "version"
	TrafficName   = "traffic"
	StatusName    = "status"
	RouteName     = "route"
)

var (
	safeRoute = common.NewSafe()
	routeData = []Route{
		{EntryId: 1, RouteId: 1, Status: "active", Name: "host", Traffic: "egress", Authority: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, RouteId: 2, Status: "active", Name: "host", Traffic: "egress", Authority: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Route - need to find some way to avoid duplicate processing, as host will start and stop due to
// Kubernetes scaling
type Route struct {
	EntryId   int       `json:"entry-id"`
	RouteId   int       `json:"route-id"`
	Status    string    `json:"status"` // Active - inactive
	CreatedTS time.Time `json:"created-ts"`

	Traffic   string `json:"traffic"` // Ingress or egress
	Name      string `json:"route"`
	Authority string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	Version   string `json:"version"`   // Semantic versioning: 2.1.0

	Timeout      int  `json:"timeout"`
	RateLimiting bool `json:"rate-limiting"`

	// Host and primary configuration are as follows:
	// Primary - route to host unless upstream failure rate exceeds a threshold, then route to secondaries
	//           based on filters. If no filters are configured then stay on primary. Once the rate falls
	//           below the threshold, fail back to the host
	// Default - route to host on startup, then rely on secondaries for all routing. Without secondary
	//           filters, remain on host.
	//
	Host    string `json:"host"`
	Primary bool   `json:"primary"`

	// Filters are only configured for egress traffic. Local is valid as is *. Blank does not include.
	RegionFilter  string `json:"region-filter"`
	ZoneFilter    string `json:"zone-filter"`
	SubZoneFilter string `json:"sub-zone-filter"`
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
		case StatusName:
			e.Status = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case AuthorityName:
			e.Authority = values[i].(string)
		case TrafficName:
			e.Traffic = values[i].(string)
		case VersionName:
			e.Version = values[i].(string)
		case RouteName:
			e.Name = values[i].(string)

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
		e.Status,
		e.CreatedTS,

		e.Traffic,
		e.Authority,
		e.Version,
		e.Name,
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
