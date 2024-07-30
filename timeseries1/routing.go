package timeseries1

type Routing struct {
	Url     string `json:"url"`
	Host    string `json:"host"`
	To      string `json:"to"` // Primary - secondary
	Percent int    `json:"percent"`
	Code    string `json:"code"`
}
