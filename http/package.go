package http

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"strings"
)

const (
	ModulePath = "github/advanced-go/observation"
	PkgPath    = ModulePath + "/http"
)

// https://localhost:8081/github/advanced-go/guidance:v1/search?q=golang

func Exchange(r *http.Request) (*http.Response, *core.Status) {
	_, status := httpx.ValidateRequest(r, ModulePath)
	if !status.OK() {
		return httpx.NewErrorResponse(status), status
	}
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:

	default:
		//_, status = activity.PostEntry[*http.Request](r.Header, r.Method, r.URL.Query(), r)
		//httpx.WriteResponse[E](w, nil, status.HttpCode(), nil)
		//return status
	}
	return nil, nil
}
