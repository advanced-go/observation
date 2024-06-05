package timeseries1

import (
	"context"
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleExchange_GetAll() {
	url, _ := url.Parse("https://www.google.search/search?region=*")
	docs1, h2, status1 := get[core.Output](context.Background(), nil, url.Query())
	fmt.Printf("test: get() -> [status:%v] [header:%v] [count:%v]\n", status1, h2, len(docs1))

	//Output:
	//test: get() -> [status:OK] [count:2]

}
