package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
	"time"
)

const (
	PkgPath   = "github/advanced-go/observation/timeseries1"
	RouteName = "timeseries-access"
	hostKey   = "host-key"
)

var resolver = uri.NewResolver([]uri.HostEntry{{Key: hostKey, Host: "www.observation.com", Proxy: false}})

// EgressRoute - upstream egress traffic route configuration
func EgressRoute(routeName string) (*controller.Config, bool) {
	switch routeName {
	case RouteName:
		return &controller.Config{RouteName: RouteName, Host: resolver.Host(hostKey), Authority: module.TimeseriesAuthority, LivenessPath: core.HealthLivenessPath, Duration: time.Second * 2}, true
	default:
		return nil, false
	}
}

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Entry, http.Header, *core.Status) {
	return get[core.Log](ctx, core.AddRequestId(h), values)
}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}
