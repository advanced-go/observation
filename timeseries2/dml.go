package timeseries2

const (
	accessLogResource = "access-log"
	accessLogSelect   = "select * from access-log"
	accessLogInsert   = "insert into access-log"

	startTimeName = "start_time"
	durationName  = "duration_ms"
	trafficName   = "traffic"
	createdTSName = "created_ts"

	regionName     = "region"
	zoneName       = "zone"
	subZoneName    = "sub_zone"
	hostName       = "host"
	instanceIdName = "instance_id"

	requestIdName = "request_id"
	relatesToName = "relates_to"
	protocolName  = "protocol"
	methodName    = "method"
	fromName      = "from"
	toName        = "to"
	uriName       = "url"
	pathName      = "path"

	statusCodeName = "status_code"
	encodingName   = "encoding"
	bytesName      = "bytes"

	routeName        = "route"
	routeToName      = "route_to"
	routePercentName = "route_percent"
	routeCodeName    = "rc"

	timeoutName        = "timeout"
	rateLimitName      = "rate_limit"
	rateBurstName      = "rate_burst"
	controllerCodeName = "cc"
)
