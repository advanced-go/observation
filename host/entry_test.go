package host

import (
	"fmt"
	"github.com/advanced-go/stdlib/json"
)

var list = []Entry{
	{"us-west", "oregon", "dc1", "www.test-host.com"},
	{"us-west", "oregon", "dc1", "localhost:8081"},
}

func ExampleEntry() {
	buf, status := json.Marshal(list)
	if !status.OK() {
		fmt.Printf("test: Entry{} -> [status:%v]\n", status)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}
