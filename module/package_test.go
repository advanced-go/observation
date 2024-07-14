package module

import (
	"fmt"
	"log"
	"runtime/debug"
)

func ExampleDebug() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}

	for _, dep := range bi.Deps {
		fmt.Printf("Dep: %+v\n", dep)
	}

	//Output:
	//fail
}
