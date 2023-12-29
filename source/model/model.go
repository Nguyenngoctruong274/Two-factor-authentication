package model

import (
	"io"
	"net/http"
	"time"
)

type Option struct {
	Method  string
	Params  string
	Body    io.Reader
	Header  http.Header
	Timeout time.Duration
	Must    int
}
type Response struct {
	PrimaryIndex  int
	StatusCode    int
	Header        http.Header
	Body          []byte
	ContentLength int64
}
