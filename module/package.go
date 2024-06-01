package module

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	Authority          = "github/advanced-go/observation"
	Name               = "observation"
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
	DocumentsAuthority      = "github/advanced-go/documents"
	DocumentsResource       = "timeseries"
	DocumentsPath           = "/github/advanced-go/documents:%stimeseries"
	DocumentsV1             = "v1"
	DocumentsV2             = "v2"
	DocumentsControllerName = "documents"
)

// config - upstream egress traffic controller configuration
var (
	config = []controller.Config{
		{DocumentsControllerName, "localhost:8081", DocumentsAuthority, core.HealthLivenessPath, time.Second * 2},
	}
)

// ControllerConfig - get the controller configuration
func ControllerConfig(ctrlName string) (controller.Config, bool) {
	return controller.GetConfig(ctrlName, config)
}
