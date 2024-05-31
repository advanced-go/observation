package timeseries1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	PkgPath = "github/advanced-go/observation/resiliency1"

	DocumentsControllerName = "documents"
)

// Controllers - egress traffic controllers
var (
	Controllers = []controller.Config{
		{DocumentsControllerName, "localhost:8081", module.DocumentsAuthority, core.HealthLivenessPath, time.Second * 2},
	}
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, url *url.URL) ([]Entry, *core.Status) {
	if url == nil || !strings.HasPrefix(url.Path, module.TimeseriesResource) {
		return nil, core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid or nil URL")))
	}
	if url.Query() == nil {
		return nil, core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("query arguments are nil")))
	}
	switch url.Path {
	case module.TimeseriesResource:
		return get[core.Log](ctx, core.AddRequestId(h), url.Query())
	default:
		return nil, core.StatusBadRequest()
	}
}

// Put - resource PUT
func Put(r *http.Request, body []Entry) *core.Status {
	if r == nil || r.URL == nil || !strings.HasPrefix(r.URL.Path, module.TimeseriesResource) {
		return core.NewStatusError(http.StatusBadRequest, errors.New("invalid URL"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status, core.RequestId(r.Header))
			return status
		}
		body = content
	}
	switch r.URL.Path {
	case module.TimeseriesResource:
		return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
	default:
		return core.StatusBadRequest()
	}
}
