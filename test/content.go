package test

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"io"
	"testing"
)

type DeserializeStatus[T any] struct {
	Error  bool
	Data   T
	Status *core.Status
}

/*
type ContentStatus[T any] struct {
	Got  DeserializeStatus[T]
	Want DeserializeStatus[T]
}

func (c ContentStatus[T]) GotCode() int {
	return c.Got.Status.Code
}

func (c ContentStatus[T]) GotItems() []T {
	return c.Got.Items
}

func (c ContentStatus[T]) WantCode() int {
	return c.Want.Status.Code
}

func (c ContentStatus[T]) WantItems() []T {
	return c.Want.Items
}


*/

func Deserialize2[T any](gotBody, wantBody io.Reader, t *testing.T) (gotStatus, wantStatus DeserializeStatus[T], success bool) {
	results(gotBody, &gotStatus)
	if gotStatus.Error {
		t.Errorf("Deserialize() %v err = %v", "got", gotStatus.Status.Err)
		return
	}
	results(wantBody, &wantStatus)
	if wantStatus.Error {
		t.Errorf("Deserialize() %v err =%v", "want", wantStatus.Status.Err)
		return
	}
	if gotStatus.Status.Code != wantStatus.Status.Code {
		t.Errorf("Deserialize() got status code = %v, want status code = %v", gotStatus.Status.Code, wantStatus.Status.Code)
		return
	}
	return gotStatus, wantStatus, true
}

func results[T any](body io.Reader, r *DeserializeStatus[T]) (bool, *core.Status) {
	r.Data, r.Status = httpx.Content[T](body)
	if !r.Status.OK() && !r.Status.NotFound() {
		r.Error = true
	}
	return r.Error, r.Status
}
