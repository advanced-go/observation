package routing1

import (
	"fmt"
	json2 "github.com/advanced-go/stdlib/json"
)

var (
	case1 = Case{
		Desc: "client: [match] cloud: []",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www/google.com",
			Route:     "google-search",
		},
			Timeout: 0,
		},
		Cloud: Cloud{RateLimiting: false},
	}

	case2 = Case{
		Desc: "client: [match,timeout] cloud: [rate-limiting]",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www/google.com",
			Route:     "google-search",
		},
			Timeout: 3000,
		},
		Cloud: Cloud{RateLimiting: true},
	}

	case3 = Case{
		Desc: "client: [match,timeout] cloud: [rate-limiting]",
		Client: Client{Match: Matcher{
			Path:      "/test/search?q=golang",
			Template:  "{path}",
			Authority: "https://www/google.com",
			Route:     "google-search",
		},
			Timeout: 3000,
		},
		Cloud: Cloud{RateLimiting: true},
	}
)

func ExampleCase_1() {
	buf, status := json2.Marshal(&case1)
	fmt.Printf("test: Case_1() -> [status:%v] [buf:%v]\n", status, string(buf))

	//Output:
	//fail
}

func ExampleCase_2() {
	buf, status := json2.Marshal(&case2)
	fmt.Printf("test: Case_2() -> [status:%v] [buf:%v]\n", status, string(buf))

	//Output:
	//fail
}
