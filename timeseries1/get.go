package timeseries1

import (
	"context"
	"github.com/advanced-go/observation/common"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
	"net/url"
)

func testOverride(ctx context.Context, resource string) context.Context {
	ex := core.ExchangeOverrideFromContext(ctx)
	if ex != nil {
		return ctx
	}
	rsc := testrsc.TS1EgressEntryTest
	if resource == IngressResource {
		rsc = testrsc.TS1IngressEntryTest
	}
	return core.NewExchangeOverrideContext(ctx, core.NewExchangeOverride("", rsc, ""))
}

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource string, values url.Values) (entries []T, h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeText)
	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	// Testing only
	ctx = testOverride(ctx, resource)

	// Set XFrom so that PostgreSQL logging is correct.
	h = httpx.SetHeader(h, core.XFrom, module.Authority)
	entries, status = pgxsql.QueryT[T](ctx, h, common.AccessLogResource, common.AccessLogSelect, values)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
		return nil, h2, status
	}
	if values != nil && len(values) > 0 {
		entries = filter[T](entries, values)
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
		h2.Set(httpx.ContentType, httpx.ContentTypeJson)
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
