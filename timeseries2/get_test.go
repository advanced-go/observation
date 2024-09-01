package timeseries2

import (
	"fmt"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGet() {
	ctx := core.NewUrlContext(nil, testrsc.TS2IngressEntryURL)
	values := make(url.Values)

	values.Add(core.RegionKey, "us-west")
	entries, _, status := get[core.Output, Entry](ctx, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Add(core.SubZoneKey, "dc1")
	entries, _, status = get[core.Output, Entry](ctx, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
