package timeseries1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) *core.Status {
	return core.StatusOK()

}
