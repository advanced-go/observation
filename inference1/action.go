package inference1

import "time"

// Action - host
type Action struct {
	// Origin + route for uniqueness
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	SubZone   string    `json:"sub-zone"`
	Host      string    `json:"host"`
	RouteName string    `json:"route"`
	AgentId   string    `json:"agent-id"`
	CreatedTS time.Time `json:"created-ts"`

	Action     string  `json:"action"`
	Timeout    int     `json:"timeout"`
	RateLimit  float64 `json:"rate-limit"`
	RateBurst  int     `json:"rate-burst"`
	Primary    string  `json:"primary"`
	Secondary  string  `json:"secondary"`
	Percentage int     `json:"percentage"`
	Status     string  `json:"status"`
}
