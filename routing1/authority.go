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
)

var (
	safeAuthority = common.NewSafe()
	authorityData = []Authority{
		{EntryId: 1, AuthorityId: 1, Status: "active", Traffic: "egress", Tag: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
		{EntryId: 1, AuthorityId: 2, Status: "active", Traffic: "egress", Tag: "github/advanced-go/observation", Version: "2.1.0", CreatedTS: time.Date(2024, 6, 10, 7, 120, 35, 0, time.UTC)},
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
	Tag     string `json:"tag"`     // github/advanced-go/observation: provider/account/repository
	Version string `json:"version"` // Semantic versioning: 2.1.0

	// Origins are only configured for egress traffic, and is used to limit host selection.
	// How to configure a primary??
	Include core.Origin `json:"include"`
	Exclude core.Origin `json:"exclude"`
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
