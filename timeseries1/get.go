package timeseries1

import (
	"context"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if query == nil {
		query = testQuery[T] //pgxsql.QueryT[T]
	}
	if h != nil {
		h.Set(core.XFrom, module.Authority)
	}
	entries, status = query(ctx, h, accessLogResource, accessLogSelect, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

func testQuery[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, _, _ string, _ map[string][]string, _ ...any) (entries []T, status *core.Status) {
	status = core.StatusOK()
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p = append(*p, entryData...)
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}
