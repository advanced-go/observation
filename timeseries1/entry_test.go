package timeseries1

import (
	"encoding/json"
	"fmt"
	"github.com/advanced-go/stdlib/access"
	"time"
)

var list = []Entry{
	{time.Now().UTC(), 100, access.EgressTraffic, "us-west", "oregon", "dc1", "www.test-host.com", "123456", "req-id", "relate-to", "HTTP/1.1", "GET", "www.google.com", "https://www.google.com/search?q-golang", "/search", 200, "gzip", 12345, "google-search", "primary", 0, ""},
	{time.Now().UTC(), 100, access.IngressTraffic, "us-west", "oregon", "dc1", "localhost:8081", "123456", "req-id", "relate-to", "HTTP/1.1", "GET", "github/advanced-go/search", "http://localhost:8081/advanced-go/search:google?q-golang", "/search", 200, "gzip", 12345, "search", "primary", 0, ""},
}

func ExampleEntry() {
	buf, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("test: Entry{} -> [err:%v]\n", err)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}
