package timeseries1

import (
	"context"
	"errors"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"net/http"
	"net/url"
)

const (
	PkgPath = "github/advanced-go/observation/timeseries1"
)

func errorInvalidURL(path string) *core.Status {
	return core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("invalid argument: URL path is invalid %v", path)))
}

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) ([]Entry, http.Header, *core.Status) {
	return get[core.Log](ctx, core.AddRequestId(h), values)

}

// Put - resource PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil || r.URL == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid URL"))
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
