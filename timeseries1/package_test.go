package timeseries1

import (
	"context"
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func ExampleExchange_GetAll() {
	url, _ := url.Parse("https://www.google.search/search?region=*")
	h := make(http.Header)
	h.Add(core.XAuthority, module.Authority)
	docs1, h2, status1 := get[core.Output](context.Background(), h, url.Query())
	fmt.Printf("test: get() -> [status:%v] [header:%v] [count:%v]\n", status1, h2, len(docs1))

	//Output:
	//test: get() -> [status:OK] [header:map[]] [count:2]

}
