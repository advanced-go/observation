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

	TS1EgressEntryURL  = TS1BasePath + "/egress-entry.json"
	TS1IngressEntryURL = TS1BasePath + "/ingress-entry.json"

	TS1GetReqURL  = TS1BasePath + "/get-req.txt"
	TS1GetRespURL = TS1BasePath + "/get-resp.txt"

	TS2IngressEntryURL = TS2BasePath + "/ingress-entry.json"
	TS2EgressEntryURL  = TS1BasePath + "/egress-entry.json"
)
