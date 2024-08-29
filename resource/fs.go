package resource

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

	Timeseries1EntryURL = TS1BasePath + "/entry.json"

	Timeseries2EntryURL = TS2BasePath + "/entry.json"
)
