package assignment1

import (
	"context"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values, resource, template string, query pgxsql.QueryFuncT[T]) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if query == nil {
		query = testQuery[T] //pgxsql.QueryT[T]
	}
	h2 = httpx.Forward(h2, h)
	h2.Set(core.XFrom, module.Authority)
	entries, status = query(ctx, h, resource, template, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}

func testQuery[T pgxsql.Scanner[T]](_ context.Context, _ http.Header, _, _ string, _ map[string][]string, _ ...any) (entries []T, status *core.Status) {
	status = core.StatusOK()
	switch p := any(&entries).(type) {
	case *[]Entry:
		*p = entryData
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	return
}