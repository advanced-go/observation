package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, u *url.URL) (entries []Entry, status *core.Status) {
	var e E
	if u == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("invalid argument: URL is nil"))
	}

	url := uri.Expansion("", module.DocumentsPath, module.DocumentsV1, u.Query())
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	httpx.Forward(req.Header, h)
	resp, status1 := httpx.DoExchange(req)
	if !status1.OK() {
		e.Handle(status1, core.RequestId(h))
		return nil, status1
	}
	entries, status = json.New[[]Entry](resp.Body, h)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
