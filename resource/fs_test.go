package resource

import (
	"fmt"
	"github.com/advanced-go/stdlib/io"
)

func ExampleReadFile() {
	name := "file:///f:/files/timeseries1/entry.json"
	bytes, status := io.ReadFile(name)
	fmt.Printf("test: ReadFile() -> [buff:%v] [status:%v]\n", len(bytes), status)

	//Output:
	//test: ReadFile() -> [buff:1609] [status:OK]

}
