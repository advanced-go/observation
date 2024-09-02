package test

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"io"
	"testing"
)

type ContentResults[T any] struct {
	Error  bool
	Items  []T
	Status *core.Status
}

type ContentStatus[T any] struct {
	Got  ContentResults[T]
	Want ContentResults[T]
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

func content[T any](gotBody, wantBody io.Reader, t *testing.T) (ContentStatus[T], bool) {
	s := ContentStatus[T]{}
	results(gotBody, &s.Got)
	if s.Got.Error {
		t.Errorf("content() %v err = %v", "got", s.Got.Status.Err)
		return s, false
	}
	results(wantBody, &s.Want)
	if s.Want.Error {
		t.Errorf("content() %v err =%v", "want", s.Want.Status.Err)
		return s, false
	}
	if s.GotCode() != s.WantCode() {
		t.Errorf("content() got status code = %v, want status code = %v", s.GotCode(), s.WantCode())
		return s, false
	}
	return s, true
}

func results[T any](body io.Reader, r *ContentResults[T]) {
	r.Items, r.Status = httpx.Content[[]T](body)
	if !r.Status.OK() && !r.Status.NotFound() {
		r.Error = true
	}
}
