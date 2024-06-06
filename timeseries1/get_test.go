package timeseries1

import (
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

const (
	getAllReq = "file://[cwd]/timeseries1test/get-all-resp-v1.txt"
)

func ExampleGet() {
	values := make(url.Values)
	h := make(http.Header)
	h.Add(BuildPath(module.TimeseriesAuthority, module.TimeseriesV1, module.TimeseriesAccessResource, nil), getAllReq)
	entries, _, status := get[core.Output](nil, h, values)

	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]

}
