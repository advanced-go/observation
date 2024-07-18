package access1

type Routing struct {
	Url            string `json:"url"`
	Host           string `json:"host"`
	RouteTo        string `json:"route-to"`
	RoutingPercent int    `json:"routing-percent"`
	RoutingCode    string `json:"routing-code"`
}
