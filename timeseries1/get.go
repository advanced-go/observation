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
	// Set XFrom so that pgxsql logging is correct.
	if h == nil {
		h = make(http.Header)
	}
	h.Set(core.XFrom, module.Authority)
	entries, status = pgxsql.QueryT[T](ctx, h, common.AccessLogResource, common.AccessLogSelect, values)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	if values != nil && len(values) > 0 {
		entries = filter[T](entries, values)
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

func filter[T pgxsql.Scanner[T]](entries []T, values url.Values) (result []T) {
	match := core.NewOrigin(values)
	customer := values.Get("customer")
	switch p := any(&result).(type) {
	case *[]Entry:
		if p != nil {
		}
		if entries2, ok := any(entries).([]Entry); ok {
			for _, e := range entries2 {
				if customer != "" && customer != e.CustomerId {
					continue
				}
				if core.OriginMatch(e.Origin(), match) {
					*p = append(*p, e)
				}
			}
		}
	default:
	}
	return result
}
