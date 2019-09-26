package http

import (
	"context"
	"net/http"
)

type Server struct {
	ctx     context.Context
	dec     DecodeReq
	before  []RequestMiddleWare
	handler Handler
	after   []ResponseMiddleWare
}

type ErrorMsg struct {
	Reason string
	code   int
}

// get req go struct from url,query,body
type DecodeReq func(ctx context.Context, req *http.Request) (interface{}, error)

// get from req , and set somethin in context
type RequestMiddleWare func(ctx context.Context, req *http.Request) context.Context

// Endpoint
type Handler func(ctx context.Context, request interface{}) (response interface{}, err error)

type ResponseMiddleWare func(ctx context.Context, resp *http.Response) context.Context

type MiddleWareOption func(server *Server)

// MiddleBefore ...
func MiddleBefore(before ...RequestMiddleWare) MiddleWareOption {
	return func(server *Server) {
		server.before = before
	}
}

// MiddleAfter ...
func MiddleAfter(after ...ResponseMiddleWare) MiddleWareOption {
	return func(server *Server) {
		server.after = after
	}
}
