package threshold1

import "time"

type Entry struct {
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	Route     string    `json:"route"`
	CreatedTS time.Time `json:"created-ts"`

	Percent int16 `json:"percent"` // Used for latency, traffic, status codes, counter, profile
	Value   int16 `json:"value"`   // Used for latency, saturation duration or traffic
	Minimum int16 `json:"minimum"` // Used for status codes to attenuate underflow, applied to the window interval
}
