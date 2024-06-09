package module

const (
	Authority             = "github/advanced-go/observation"
	RouteName             = "observation"
	Version               = "2.2.2"
	Ver1                  = "v1"
	Ver2                  = "v2"
	ObservationTimeseries = "timeseries"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)

// Upstream authorities/resources
const (
	TimeseriesAuthority        = "github/advanced-go/timeseries"
	TimeseriesAccessResourceV1 = "v1/access"
)
