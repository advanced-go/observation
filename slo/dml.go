package slov

const (
	customerIdName     = "customer_id"
	startTimeName      = "start_time"
	durationName       = "duration_ms"
	durationStrName    = "duration_str"
	trafficName        = "traffic"
	regionName         = "region"
	zoneName           = "zone"
	subZoneName        = "sub_zone"
	serviceName        = "service"
	instanceIdName     = "instance_id"
	routeNameName      = "route_name"
	requestIdName      = "request_id"
	urlName            = "url"
	protocolName       = "protocol"
	methodName         = "method"
	hostName           = "host"
	pathName           = "path"
	statusCodeName     = "status_code"
	bytesSentName      = "bytes_sent"
	statusFlagsName    = "status_flags"
	timeoutName        = "timeout"
	rateLimitName      = "rate_limit"
	rateBurstName      = "rate_burst"
	retryName          = "retry"
	retryRateLimitName = "retry_rate_limit"
	retryRateBurstName = "retry_rate_burst"
	failoverName       = "failover"

	//entrySelect = "SELECT * FROM slo_entry order by start_time limit 2"
	entrySelect = "SELECT region,customer_id,start_time,duration_str,traffic,rate_limit FROM slo_entry order by start_time desc limit 2"

	entryInsert = "INSERT INTO slo_entry (" +
		"customer_id,start_time,duration_ms,duration_str,traffic," +
		"region,zone,sub_zone,service,instance_id,route_name," +
		"request_id,url,protocol,method,host,path,status_code,bytes_sent,status_flags," +
		"timeout,rate_limit,rate_burst,retry,retry_rate_limit,retry_rate_burst,failover) VALUES"
)
