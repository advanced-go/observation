package test

import (
	http2 "github.com/advanced-go/observation/http"
	"github.com/advanced-go/observation/testrsc"
	"github.com/advanced-go/observation/timeseries1"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"io"
	"net/http"
	"reflect"
	"testing"
)

type contentResults struct {
	error  bool
	entry  []timeseries1.Entry
	status *core.Status
}

type contentStatus struct {
	got  contentResults
	want contentResults
}

func (c contentStatus) gotCode() int {
	return c.got.status.Code
}

func (c contentStatus) gotEntry() []timeseries1.Entry {
	return c.got.entry
}

func (c contentStatus) wantCode() int {
	return c.want.status.Code
}

func (c contentStatus) wantEntry() []timeseries1.Entry {
	return c.want.entry
}

func TestExchange1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		//{name: "read-request-error", req: readRequest("", t), resp: readResponse(testrsc.TS1GetRespURL, t), status: core.StatusOK()},
		{name: "get-entry", req: readRequest(testrsc.TS1GetReqURL, t), resp: readResponse(testrsc.TS1GetRespURL, t), status: core.StatusOK()},
	}
	for _, tt := range tests {
		cont := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, %v", status.Code, tt.status.Code, status.Err)
				cont = false
			}
			if cont && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				cont = false
			}
			cs := contentStatus{}
			if cont {
				cs, cont = content(resp.Body, tt.resp.Body, t)
			}
			if cont {
				if !reflect.DeepEqual(cs.gotEntry(), cs.wantEntry()) {
					t.Errorf("Exchange() got = %v, want %v", cs.gotEntry(), cs.wantEntry())
				}
			}
		})
	}
}

func content(gotBody, wantBody io.Reader, t *testing.T) (contentStatus, bool) {
	s := contentStatus{}
	results(gotBody, &s.got, t)
	if s.got.error {
		t.Errorf("content() %v err : %v", "got", s.got.status.Err)
		return s, false
	}
	results(wantBody, &s.want, t)
	if s.want.error {
		t.Errorf("content() %v err : %v", "want", s.want.status.Err)
		return s, false
	}
	if s.gotCode() != s.wantCode() {
		t.Errorf("content() got status code : %v, want status code : %v", s.gotCode(), s.wantCode())
		return s, false
	}
	return s, true
}

func results(body io.Reader, r *contentResults, t *testing.T) {
	r.entry, r.status = httpx.Content[[]timeseries1.Entry](body)
	if !r.status.OK() && !r.status.NotFound() {
		r.error = true
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
