package http

import (
	"fmt"
	"github.com/advanced-go/observation/timeseries1"
	"github.com/advanced-go/stdlib/json"
	"net/http"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ := http.NewRequest("", "http://www.google.com/search?q=golang", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ = http.NewRequest("", "http://www.google.com/github/advanced-go/observation", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	//Output:
	//test: Exchange(nil) -> [status:Invalid Argument [error: request is nil]] [status-code:500]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, authority does not match: "/search" "github/advanced-go/observation"]] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, path only contains an authority: "/github/advanced-go/observation"]] [status-code:400]

}

func ExampleExchange_Timeseries() {
	uri := "http://localhost:8081/github/advanced-go/observation:v1/timeseries?region=*"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := Exchange(req)
	if !status.OK() {
		fmt.Printf("test: Exchange() -> [status:%v]\n", status)
	} else {
		entries, status1 := json.New[[]timeseries1.Entry](resp.Body, resp.Header)
		fmt.Printf("test: Exchange() -> [status:%v] [status-code:%v] [bytes:%v] [content:%v]\n", status1, resp.StatusCode, resp.ContentLength, entries)
	}

	//Output:
	//test: Exchange() -> [status:OK] [status-code:200] [bytes:359] [content:[{region1 Zone1  www.host1.com active   100ms 125 25} {region1 Zone2  www.host2.com inactive   250ms 100 10}]]

}
