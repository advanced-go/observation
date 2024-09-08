package timeseries1

import (
	"fmt"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
)

func ExamplePut() {
	entries, _ := json.New[[]Entry](testrsc.TS1EgressEntry, nil)

	ex := core.NewExchangeOverride("", "", json.StatusTimeoutUri)
	ctx := core.NewExchangeOverrideContext(nil, ex)

	_, status := put[core.Output, Entry](ctx, nil, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
