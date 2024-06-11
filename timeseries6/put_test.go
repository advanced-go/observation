package timeseries6

import (
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"time"
)

const (
	putResp = "file://[cwd]/timeseries1test/put-resp-v1.txt"
)

func ExamplePut() {
	h := make(http.Header)
	h.Add(uri.BuildPath(module.TimeseriesAuthority, module.TimeseriesAccessResourceV1, nil), putResp)

	_, status := put[core.Output](nil, h, nil)
	fmt.Printf("test: put(nil,h,nil) -> [status:%v]\n", status)

	_, status = put[core.Output](nil, h, []Entry{{StartTime: time.Now().UTC()}})
	fmt.Printf("test: put(nil,h,[]Entry) -> [status:%v]\n", status)

	//Output:
	//test: put(nil,h,nil) -> [status:OK]
	//test: put(nil,h,[]Entry) -> [status:Timeout]

}
