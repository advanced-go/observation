package timeseries1

import (
	"bytes"
	"context"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"io"
	"net/http"
)

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (http.Header, *core.Status) {
	var e E

	url := uri.Expansion("", module.TimeseriesPath, module.TimeseriesV1, nil)
	rc, _, status := createReadCloser(body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, status
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, url, rc)
	httpx.Forward(req.Header, h, core.XAuthority)
	_, status = httpx.DoExchange(req)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return nil, status
}

func createReadCloser(body any) (io.ReadCloser, int64, *core.Status) {
	switch ptr := body.(type) {
	case []Entry:
		return json2.NewReadCloser(body)
	case []byte:
		return io.NopCloser(bytes.NewReader(ptr)), int64(len(ptr)), core.StatusOK()
	default:
		return nil, 0, core.NewStatusError(core.StatusInvalidArgument, core.NewInvalidBodyTypeError(body))
	}
}
