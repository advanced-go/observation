package slov

type GetConstraints interface {
	[]byte | []Entry
}

const (
	pkgPath = "github/advanced-go/observation"
)

var (
	getLoc = pkgPath + "/get"
)

/*

func Get[E core.ErrorHandler, T GetConstraints](ctx context.Context, url *url.URL) (T, *core.Status) {
	var e E
	var t T

	//rows, status := pgxsql.QueryT[E](ctx, pgxsql.NewQueryRequest(entryResource, entrySelect, nil))
	//if !status.OK() {
	//	return nil, status
	//}
	//events, err := pgxsql.Scan[Entry](rows)
	//if err != nil {
	//	return nil, e.HandleWithContext(ctx, getLoc, err)
	//}
	switch ptr := any(&t).(type) {
	case *[]byte:
		buf, err1 := json.Marshal(events)
		if err1 != nil {
			return nil, e.HandleWithContext(ctx, getLoc, err1).SetCode(runtime.StatusInvalidContent)
		}
		*ptr = buf
	case *[]Entry:
		*ptr = events
	}
	return t, runtime.NewStatusOK()
}

func Ping[E template.ErrorHandler](ctx context.Context) *runtime.Status {
	return pgxsql.Ping[E](ctx)
}


*/
