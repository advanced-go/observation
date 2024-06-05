package module

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"time"
)

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
	TimeseriesAuthority = "github/advanced-go/timeseries"
	TimeseriesPath      = TimeseriesAuthority + ":%s" + TimeseriesAccessResource
	TimeseriesV1        = "v1"
	//TimeseriesV2        = "v2"
	TimeseriesRouteName      = "timeseries-access"
	TimeseriesAccessResource = "access"
)

// Routes - upstream egress traffic route configuration
var (
	Routes = []controller.Config{
		{TimeseriesRouteName, "localhost:8081", TimeseriesAuthority, core.HealthLivenessPath, time.Second * 2},
	}
)

// GetRoute - get the route configuration
func GetRoute(routeName string) (controller.Config, bool) {
	return controller.GetRoute(routeName, Routes)
}
