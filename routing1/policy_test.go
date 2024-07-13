package routing1

import "fmt"

func ExampleLastPolicy() {

	fmt.Printf("test: lastPolicy() -> [entry:%v]\n", lastPolicy().EntryId)

	//Output:
	//test: lastPolicy() -> [entry:1]

}
