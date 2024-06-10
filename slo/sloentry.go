package slov

const (
	entryResource = "slo-entryv1"
)

// Thresholds - common parameters
type Thresholds struct {
	Watch   int16 // Range 1 - 99
	Percent int16 // Used for latency, traffic, status codes, counter, profile
	Value   int16 // Used for latency, saturation duration or traffic
	Minimum int16 // Used for status codes to attenuate underflow, applied to the window interval
}

type Entry struct {
	// Identification
	Id         int32
	CustomerId int32
	Name       string // SLO name, unique within customer, service name and route
	Service    string
	RouteName  string
	Disabled   bool // Whether to process or not

	// Selection
	TrafficType       int32  // Traffic type - ingress, egress, ping, profile, counter
	TrafficProtocol   int32  // Traffic protocol - gRPC, HTTP10, HTTP11, HTTP2, HTTP3
	FilterStatusCodes string // Used as a configurable filter

	// SCF = status code filter  SCL - status code list
	// Category -
	//  latency       - SC Filter: [..], percentile: 1-99, latency value: 500ms
	//  status codes  - SC Filter: [..], minimum: value or RPS, SC List: [..], percentage of traffic
	//  traffic       - SC Filter: [..], calculated historical value, comparison: [..], RPS value
	//  saturation    - Ping traffic, percentile: 1-99, latency value: 500ms
	//
	//  saturation-metrics
	//  counter
	//  profile
	// TODO: Saturation SLOs need separate processing, less adaptive and more often. These are early warning
	Category int16

	// Thresholds
	Thresholds Thresholds

	// Computation inputs
	// Traffic - this is current relative to a historical value
	// TODO : is previous days an average, or a comparison against each individual day?
	RPSLowComparison int16 // Values : None, Previous N Days, Same Day Last Week
	// The comparison for high would be a percentage of the metric
	RPSHighComparison int16  // Values : None, Previous N Days, Same Day Last Week, All-time
	StatusCodes       string // Comma seperated list of status codes, for a status code SLO

	// Interval range and window size
	// TODO: let adaptive determine optimal processing and window intervals
	//       using the configured value as being associated with the from interval
	//       The adaptive scaling needs to be logarithmic or exponential?
	//       The window interval should be larger than the from interval as this would catch issues
	//       that began at the end of a previous timeframe and ended at the beginning of the next window
	// TODO: RPS SLOs need an option for an interval of 1 day
	// TODO: Profile SLOs need to support granularity down to seconds
	ProcessingInterval int16 // Minutes
	WindowInterval     int16 // Minutes

}

func (Entry) Scan(columnNames []string, values []any) (entry Entry, err error) {
	return Entry{}, nil
}
