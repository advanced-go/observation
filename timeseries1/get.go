package timeseries1

import (
	"context"
	"github.com/advanced-go/observation/common"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	if h != nil {
		h.Set(core.XFrom, module.Authority)
	}
	entries, status = pgxsql.QueryT[T](ctx, h, common.AccessLogResource, common.AccessLogSelect, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	entries = filter[T](entries, values)
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

func filter[T pgxsql.Scanner[T]](entries []T, values url.Values) (result []T) {
	match := core.NewOrigin(values)
	switch p := any(&result).(type) {
	case *[]Entry:
		if p != nil {
		}
		if entries2, ok := any(entries).([]Entry); ok {
			for _, e := range entries2 {
				if core.OriginMatch(e.Origin(), match) {
					*p = append(*p, e)
				}
			}
		}
	default:
	}
	return result
}
