package test

import (
	http2 "github.com/advanced-go/observation/http"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/observation/timeseries1"
	"github.com/advanced-go/stdlib/core"
	httpt "github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"reflect"
	"testing"
)

func TestExchange1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		//{name: "read-request-error", req: readRequest("", t), resp: readResponse(testrsc.TS1GetRespURL, t), status: core.StatusOK()},
		{name: "get-entry", req: httpt.NewRequestTest(testrsc.TS1GetReqURL, t), resp: httpt.NewResponseTest(testrsc.TS1GetRespURL, t), status: core.StatusOK()},
	}
	for _, tt := range tests {
		cont := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				cont = false
			}
			if cont && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				cont = false
			}
			cs := ContentStatus[timeseries1.Entry]{}
			if cont {
				cs, cont = content[timeseries1.Entry](resp.Body, tt.resp.Body, t)
			}
			if cont {
				if !reflect.DeepEqual(cs.GotItems(), cs.WantItems()) {
					t.Errorf("Exchange() got = %v, want %v", cs.GotItems(), cs.WantItems())
				}
			}
		})
	}
}

/*
	if !reflect.DeepEqual(got, tt.want) {
		t.Errorf("Exchange() got = %v, want %v", got, tt.want)
	}
	if !reflect.DeepEqual(got1, tt.want1) {
		t.Errorf("Exchange() got1 = %v, want %v", got1, tt.want1)
	}

*/
