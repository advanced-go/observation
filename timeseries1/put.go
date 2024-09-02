package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/observation/common"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, body []T) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	if h != nil {
		h.Set(core.XFrom, module.Authority)
	}
	_, status = pgxsql.InsertT[T](ctx, h, common.AccessLogResource, common.AccessLogInsert, body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
