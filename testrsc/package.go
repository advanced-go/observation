package testrsc

const (
	ts1BasePath = "file:///f:/files/timeseries1"

	TS1EgressEntry      = ts1BasePath + "/egress-entry.json"
	TS1EgressEntryTest  = ts1BasePath + "/entry-test.json"
	TS1IngressEntry     = ts1BasePath + "/ingress-entry.json"
	TS1IngressEntryTest = ts1BasePath + "/ingress-entry-test.json"

	TS1GetReq  = ts1BasePath + "/get-req.txt"
	TS1GetResp = ts1BasePath + "/get-resp.txt"

	ts2BasePath         = "file:///f:/files/timeseries2"
	TS2IngressEntry     = ts2BasePath + "/ingress-entry.json"
	TS2IngressEntryTest = ts2BasePath + "/ingress-entry-test.json"
	TS2EgressEntry      = ts2BasePath + "/egress-entry.json"
	TS2EgressEntryTest  = ts2BasePath + "/entry-test.json"
)

const (
	th1BasePath = "file:///f:/files/threshold1"

	TH1EntryTest = th1BasePath + "/entry-test.json"
)
