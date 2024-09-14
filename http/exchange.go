package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/observation/threshold1"
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
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		resp1, _ := httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
		return resp1, status
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case timeseries:
		resp, status1 := timeseriesExchange[core.Log](r, p)
		resp.Header.Add(core.XRoute, timeseries1.Route)
		return resp, status1
	case threshold:
		resp, status1 := thresholdExchange[core.Log](r, p)
		resp.Header.Add(core.XRoute, threshold1.Route)
		return resp, status1
	case core.VersionPath:
		resp, status1 := httpx.NewVersionResponse(module.Version), core.StatusOK()
		return resp, status1
	case core.AuthorityPath:
		resp, status1 := authorityResponse, core.StatusOK()
		return resp, status1
	case core.HealthReadinessPath, core.HealthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, testresource not found: [%v]", p.Resource)))
		return httpx.NewResponse[core.Log](status.HttpCode(), h2, status.Err)
	}
}
