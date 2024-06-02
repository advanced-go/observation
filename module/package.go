package module

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	Authority          = "github/advanced-go/observation"
	RouteName          = "observation"
	Version            = "2.2.2"
	Ver1               = "v1"
	Ver2               = "v2"
	TimeseriesResource = "timeseries"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)

// Upstream authorities/resources
const (
	DocumentsAuthority = "github/advanced-go/documents"
	DocumentsPath      = DocumentsAuthority + ":%stimeseries"

	DocumentsV1        = "v1"
	DocumentsV2        = "v2"
	DocumentsRouteName = "documents"
	DocumentsResource  = "timeseries"
)

// Routes - upstream egress traffic route configuration
var (
	Routes = []controller.Config{
		{DocumentsRouteName, "localhost:8081", DocumentsAuthority, core.HealthLivenessPath, time.Second * 2},
	}
)

// GetRoute - get the route configuration
func GetRoute(routeName string) (controller.Config, bool) {
	return controller.GetRoute(routeName, Routes)
}
