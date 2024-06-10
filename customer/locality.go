package customer

import "strings"

const (
	UrnSeparator = ":"
)

type Locality struct {
	Region string // AWS - Region (Non AWS - Data Center)
	Zone   string // AWS - Zone or cluster name if only one cluster (Non AWS - cluster name)
	//SubZone string // AWS Only - multiple clusters in an AWS zone
}

func CreateHostUrn(region, zone string) string {
	return CreateUrn(&Locality{Region: region, Zone: zone})
}

func CreateUrn(loc *Locality) string {
	if loc == nil || loc.Region == "" {
		return ""
	}
	urn := loc.Region
	return addSegment(urn, loc.Zone)
	//return addSegment(urn, loc.SubZone)
}

func addSegment(urn, s string) string {
	if s != "" {
		urn += UrnSeparator + s
	}
	return urn
}

func CreateLocality(urn string) *Locality {
	if urn == "" {
		return nil
	}
	loc := new(Locality)
	tokens := strings.Split(urn, UrnSeparator)
	for i, t := range tokens {
		if t == "" {
			continue
		}
		switch i {
		case 0:
			loc.Region = t
		case 1:
			loc.Zone = t
			//case 2:
			//	loc.SubZone = t
		}
	}
	return loc
}
