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
	AuthorityIdName = "authority_id"
	TagName         = "tag"
	VersionName     = "version"
	TrafficName     = "traffic"
	StatusName      = "status"
	RouteName       = "route"
)

var (
	safeAuthority = common.NewSafe()
	authorityData = []Authority{
		{EntryId: 1, AuthorityId: 1, Status: "active", Route: "host", Traffic: "egress", Tag: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, AuthorityId: 2, Status: "active", Route: "host", Traffic: "egress", Tag: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
	}
)

// Authority - need to find some way to avoid duplicate processing, as host will start and stop due to
// Kubernetes scaling
type Authority struct {
	EntryId     int       `json:"entry-id"`
	AuthorityId int       `json:"authority-id"`
	Status      string    `json:"status"` // Active - inactive
	CreatedTS   time.Time `json:"created-ts"`

	Traffic string `json:"traffic"` // Ingress or egress
	Route   string `json:"route"`
	Tag     string `json:"tag"`     // github/advanced-go/observation: provider/account/repository
	Version string `json:"version"` // Semantic versioning: 2.1.0

	Timeout      int  `json:"timeout"`
	RateLimiting bool `json:"rate-limiting"`

	// Host and primary configuration are as follows:
	// Primary - route to host unless upstream failure rate exceeds a threshold, then route to secondaries
	//           based on filters. If not filters are configured then stay on primary. Once the rate falls
	//           below the threshold, fail back to the host
	// Default - route to host on startup, then rely on secondaries for all routing. Without secondary
	//           filters, then remain on host.
	//
	Host    string `json:"host"`
	Primary bool   `json:"primary"`

	// Filters are only configured for egress traffic. Local is valid as is *. Blank does not include.
	RegionFilter  string `json:"region-filter"`
	ZoneFilter    string `json:"zone-filter"`
	SubZoneFilter string `json:"sub-zone-filter"`
}

/*
func (e Authority) Origin() core.Origin {
	return core.Origin{Region: e.Region, Zone: e.Zone, SubZone: e.SubZone, Host: e.Host}
}


*/

func (Authority) Scan(columnNames []string, values []any) (e Authority, err error) {
	for i, name := range columnNames {
		switch name {
		case EntryIdName:
			e.EntryId = values[i].(int)
		case AuthorityIdName:
			e.EntryId = values[i].(int)
		case StatusName:
			e.Status = values[i].(string)
		case CreatedTSName:
			e.CreatedTS = values[i].(time.Time)

		case TagName:
			e.Tag = values[i].(string)
		case TrafficName:
			e.Traffic = values[i].(string)
		case VersionName:
			e.Version = values[i].(string)
		case RouteName:
			e.Route = values[i].(string)

		default:
			err = errors.New(fmt.Sprintf("invalid field name: %v", name))
			return
		}
	}
	return
}

func (e Authority) Values() []any {
	return []any{
		e.EntryId,
		e.AuthorityId,
		e.Status,
		e.CreatedTS,

		e.Traffic,
		e.Tag,
		e.Version,
		e.Route,
	}
}

func (Authority) Rows(entries []Entry) [][]any {
	var values [][]any

	for _, e := range entries {
		values = append(values, e.Values())
	}
	return values
}

func validAuthority(values url.Values, e Entry) bool {
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
