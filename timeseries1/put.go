package timeseries1

import (
	"context"
	"errors"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/postgresql/pgxsql"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"time"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, body []T, insert pgxsql.InsertFuncT[T]) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	if insert == nil {
		insert = testInsert[T] //pgxsql.InsertT[T]
	}
	if h != nil {
		h.Set(core.XFrom, module.Authority)
	}
	_, status = insert(ctx, h, accessLogResource, accessLogInsert, body)
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}

func testInsert[T pgxsql.Scanner[T]](ctx context.Context, h http.Header, resource, template string, entries []T, args ...any) (tag pgxsql.CommandTag, status *core.Status) {
	status = core.StatusOK()
	switch p := any(&entries).(type) {
	case *[]Entry:
		for _, e := range *p {
			e.CreatedTS = time.Now().UTC()
			entryData = append(entryData, e)
		}
	default:
		status = core.NewStatusError(http.StatusBadRequest, core.NewInvalidBodyTypeError(entries))
	}
	if status.OK() {
		tag.RowsAffected = int64(len(entries))
	}
	return

	return pgxsql.CommandTag{}, core.NewStatus(http.StatusTeapot)
}

/*
func postEntryHandler[E runtime.ErrorHandler](ctx context.Context, h http.Header, method string, body any) (any, *runtime.Status) {
	var e E

	switch strings.ToUpper(method) {
	case http.MethodPut:
		entries, status := createEntries(body)
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
			return nil, status
		}
		if len(entries) == 0 {
			status = runtime.NewStatusError(runtime.StatusInvalidContent, postEntryHandlerLoc, errors.New("error: no entries found"))
			e.Handle(status, runtime.RequestId(h), "")
			return nil, status
		}
		_, status = put(ctx, h, entries) // pgxsql.NewInsertRequest(h, lookup(rscAccessLog), accessLogInsert, entries[0].CreateInsertValues(entries)))
		if !status.OK() {
			e.Handle(status, runtime.RequestId(h), postEntryHandlerLoc)
		}
		return nil, status
	default:
		return nil, runtime.NewStatus(http.StatusMethodNotAllowed)
	}
}


*/
/*
func createEntries(body any) (entries []EntryV1, status *core.Status) {
	if body == nil {
		return nil, core.NewStatus(core.StatusInvalidContent)
	}
	switch ptr := body.(type) {
	case []EntryV1:
		entries = ptr
	case []byte:
		entries, status = json2.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	case io.ReadCloser:
		entries, status = json2.New[[]EntryV1](ptr, nil)
		if !status.OK() {
			return nil, status.AddLocation()
		}
	default:
		return nil, core.NewStatusError(runtime.StatusInvalidContent, runtime.NewInvalidBodyTypeError(body))
	}
	return entries, core.StatusOK()
}


*/
