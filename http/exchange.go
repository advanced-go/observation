package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/observation/timeseries1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

// https://localhost:8081/github/advanced-go/observation:v1/search?q=golang

var (
	authorityResponse = httpx.NewAuthorityResponse(module.Authority)
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (resp *http.Response, status *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status1 := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse[core.Log](status1.HttpCode(), h2, status1.Err)
	}
	p, status2 := httpx.ValidateURL(r.URL, module.Authority)
	if !status2.OK() {
		resp1, _ := httpx.NewResponse[core.Log](status2.HttpCode(), h2, status2.Err)
		return resp1, status2
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case module.ObservationTimeseries:
		resp, status = timeseriesExchange[core.Log](r, p)
		resp.Header.Add(core.XRoute, timeseries1.Route)
		return
	case core.VersionPath:
		resp, status = httpx.NewVersionResponse(module.Version), core.StatusOK()
		resp.Header.Add(core.XRoute, module.VersionRoute)
		return
	case core.AuthorityPath:
		resp, status = authorityResponse, core.StatusOK()
		return
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, testresource not found: [%v]", p.Resource)))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
}
