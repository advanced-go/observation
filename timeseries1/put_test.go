package timeseries1

import (
	"fmt"
	"github.com/advanced-go/stdlib/core"
	"time"
)

const (
	updateRsc = "test"
	updateSql = "UPDATE access_log"
	status504 = "file://[cwd]/timeseries1test/status-504.json"
)

var event = Entry{
	StartTime:  time.Now().UTC().AddDate(1, 2, 0), //ate(2023, 1, 1, 14, 12, 15, 251097, time.UTC),
	Duration:   450,
	Traffic:    "egress",
	Region:     "california",
	Zone:       "san francisco",
	SubZone:    "loma alta",
	InstanceId: "12345",
	Route:      "timeseries-egress",
	RequestId:  "67890",
	Url:        "urn:postgres:exec",
	Protocol:   "urn",
	Method:     "post",
	Host:       "postgres",
	Path:       "exec.",
	StatusCode: 200,
	Bytes:      -1,
	ReasonCode: "RL",
	Timeout:    500,
	RateLimit:  100,
	RateBurst:  25,
}

var event2 = Entry{
	StartTime: time.Date(2023, 2, 20, 5, 45, 12, 123456, time.UTC),
	//StartTime:      time.Now().UTC(),
	Duration:   45,
	Traffic:    "ingress",
	Region:     "nevada",
	Zone:       "las vegas",
	SubZone:    "rfd #1",
	InstanceId: "67890",
	Route:      "timeseries-ingress",
	RequestId:  "1234-5678-9012",
	Url:        "urn:postgres:exec",
	Protocol:   "urn",
	Method:     "post",
	Host:       "postgres",
	Path:       "exec.",
	StatusCode: 404,
	Bytes:      -1,
	ReasonCode: "TO",
	Timeout:    300,
	RateLimit:  45,
	RateBurst:  105,
}

func ExamplePut() {
	entries := []Entry{event, event2}

	_, status := put[core.Output](nil, nil, entries, testInsert[Entry])
	fmt.Printf("test: put(nil,nil,entries,testInsert) -> [status:%v] [entries:%v]\n", status, len(entryData))

	//Output:
	//test: put(nil,nil,entries,testInsert) -> [status:OK] [entries:6]

}
