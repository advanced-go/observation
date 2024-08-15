package access1

import "time"

// Header - entry header
type Header struct {
	Region      string    `json:"region"`
	Zone        string    `json:"zone"`
	SubZone     string    `json:"sub-zone"`
	Host        string    `json:"host"`
	InstanceId  string    `json:"instance-id"`
	CreatedTS   time.Time `json:"created-ts"`
	GuidanceKey string    `json:"guidance-key"` // How to query a guidance entry, which is a part of the host
	Status      string    `json:"status"`       // active,in-active

}
