package testrsc

import (
	"embed"
	"github.com/advanced-go/stdlib/io"
)

//go:embed files
var f embed.FS

func init() {
	io.Mount(f)
}

const (
	TS1BasePath = "file:///f:/files/timeseries1"
	TS2BasePath = "file:///f:/files/timeseries2"

	TS1EgressEntry      = TS1BasePath + "/egress-entry.json"
	TS1EgressEntryTest  = TS1BasePath + "/egress-entry-test.json"
	TS1IngressEntry     = TS1BasePath + "/ingress-entry.json"
	TS1IngressEntryTest = TS1BasePath + "/ingress-entry-test.json"

	TS1GetReq  = TS1BasePath + "/get-req.txt"
	TS1GetResp = TS1BasePath + "/get-resp.txt"

	TS2IngressEntry     = TS2BasePath + "/ingress-entry.json"
	TS2IngressEntryTest = TS1BasePath + "/ingress-entry-test.json"
	TS2EgressEntry      = TS2BasePath + "/egress-entry.json"
	TS2EgressEntryTest  = TS1BasePath + "/egress-entry-test.json"
)
