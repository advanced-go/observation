package test

import (
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"testing"
)

func readRequest(uri any, t *testing.T) *http.Request {
	req, status := httpxtest.ReadRequest(uri)
	if status.OK() {
		return req
	}
	t.Errorf("ReadRequest() err = %v", status.Err.Error())
	req2, _ := http.NewRequest("", "http://somedomain.com/invalid-uri", nil)
	return req2
}

func readResponse(uri any, t *testing.T) *http.Response {
	resp, status := httpx.ReadResponse(uri)
	if status.OK() {
		return resp
	}
	t.Errorf("ReadResponse() err = %v", status.Err.Error())
	return &http.Response{StatusCode: http.StatusTeapot}
}
