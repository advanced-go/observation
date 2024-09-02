package timeseries1

import (
	"fmt"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGet() {
	values := make(url.Values)
	ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.TS1EgressEntryURL, ""))

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
