package timeseries1

import (
	"context"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, nil, core.StatusNotFound()
	}
	url := uri.Expansion("", module.TimeseriesPath, module.TimeseriesV1, values)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	httpx.Forward(req.Header, h, core.XAuthority)
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		return nil, h2, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
