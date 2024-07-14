package routing1

import "time"

// IngressAuthorities - need to find some way to avoid duplicate processing, as host will start and stop due to
// / Updating is allowed as changes should be minimum
type IngressAuthorities struct {
	EntryId   int       `json:"entry-id"`
	RouteId   int       `json:"route-id"`
	CreatedTS time.Time `json:"created-ts"`
	UpdatedTS time.Time `json:"created-ts"`

	//Traffic      string `json:"traffic"` // Ingress or egress
	//RouteName string `json:"route"`
	Authority string `json:"authority"` // github/advanced-go/observation: provider/account/repository
	Version   string `json:"version"`   // Only for ingress
	//Timeout   int    `json:"timeout"`
	//RateLimiting bool   `json:"rate-limiting"`

	// Re-evalute if needed. Host can either be static or a preference for dynamic routing
	//Static       string `json:"static"`
	//Host string `json:"host"` // Enables static routing.

}
