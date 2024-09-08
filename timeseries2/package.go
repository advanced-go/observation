package timeseries2

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	PkgPath         = "github/advanced-go/observation/timeseries2"
	EgressResource  = "egress"
	IngressResource = "ingress"
)

// Get - timeseries2 resource GET
func Get(ctx context.Context, h http.Header, u *url.URL) (entries []Entry, h2 http.Header, status *core.Status) {
	if u == nil {
		return nil, h, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	rsc := ""
	if strings.Contains(u.Path, EgressResource) {
		rsc = EgressResource
	} else {
		if strings.Contains(u.Path, IngressResource) {
			rsc = IngressResource
		} else {
			return nil, h, core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not ingress or egress"))
		}
	}
	return get[core.Log, Entry](ctx, core.AddRequestId(h), rsc, u.Query())
}

// Put - timeseries2 PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

func IngressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func EgressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}
