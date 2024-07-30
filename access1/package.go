package access1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
)

const (
	PkgPath           = "github/advanced-go/observation/access1"
	accessLogResource = "access-log"
)

func IngressRateLimitingQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func IngressRedirectQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func EgressRateLimitingQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func EgressRedirectQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}

func EgressFailoverQuery(ctx context.Context, origin core.Origin) ([]Entry, *core.Status) {
	return nil, core.StatusOK()
}
