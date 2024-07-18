package inference1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	PkgPath           = "github/advanced-go/observation/inference1"
	inferenceResource = "inference"
)

// Get - resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, status *core.Status) {
	return []Entry{}, core.StatusOK()
}

// EgressQuery - query egress inference
func EgressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

// IngressQuery - query ingress inference
func IngressQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

/*



// IngressInsert - insert ingress inference
func IngressInsert(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}

// IngressInsertInterval - insert ingress interval inference
func IngressInsertInterval(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}


// EgressInsert - insert egress inference
func EgressInsert(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}

// EgressInsertInterval - insert egress interval inference
func EgressInsertInterval(ctx context.Context, h http.Header, e Entry) *core.Status {
	_, status := put[core.Log, Entry](ctx, core.AddRequestId(h), inferenceResource, "", []Entry{e}, nil)
	return status
}


*/
