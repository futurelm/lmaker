package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	transportHttp "github.com/lmfuture-ma/lmaker/transport/http"

	"github.com/lmfuture-ma/lmaker/example/todolist/dto"
)

type handlers struct {
	ListTodos transportHttp.Handler
	GetTodo   transportHttp.Handler
	AddTodo   transportHttp.Handler
}

func MakeHandler(service TodolistService, e *mux.Router) *mux.Router {
	handlers := registerHandler(service)
	e.Handle("/listTodos", transportHttp.NewHttpServer(DecodeListTodosReq, handlers.ListTodos)).Methods(http.MethodGet)
	e.Handle("/getTodo/{int}", transportHttp.NewHttpServer(DecodeGetTodoReq, handlers.GetTodo)).Methods(http.MethodGet)
	e.Handle("/addTodo", transportHttp.NewHttpServer(DecodeAddTodoReq, handlers.AddTodo)).Methods(http.MethodPost)
	return e
}

func registerHandler(s TodolistService) handlers {
	var handler handlers
	handler.ListTodos = MakeListTodosEndpoint(s)
	handler.GetTodo = MakeGetTodoEndpoint(s)
	handler.AddTodo = MakeAddTodoEndpoint(s)
	return handler
}
func MakeListTodosEndpoint(s TodolistService) transportHttp.Handler {
	return func(ctx context.Context, structReq interface{}) (interface{}, error) {
		todos, err := s.ListTodos(ctx)
		var rs *dto.ListTodosResponse
		if err != nil {
			rs = &dto.ListTodosResponse{}
		} else {
			rs = &dto.ListTodosResponse{Todos: todos}
		}
		return rs, err
	}
}

func DecodeListTodosReq(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
func MakeGetTodoEndpoint(s TodolistService) transportHttp.Handler {
	return func(ctx context.Context, structReq interface{}) (interface{}, error) {
		req := structReq.(*dto.GetTodoReq)
		todo, err := s.GetTodo(ctx, req.Int)
		var rs *dto.GetTodoResp
		if err != nil {
			rs = &dto.GetTodoResp{}
		} else {
			rs = &dto.GetTodoResp{Todo: todo}
		}
		return rs, err
	}
}

func DecodeGetTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	req := new(dto.GetTodoReq)
	// Order will be: URL parameters < query params < form data < JSON body
	// DEPRECATED: it is recommended to set 'http-parameter-priority: url' in grabkit.yaml
	if err := transportHttp.DecodeMuxVars(r, req); err != nil {
		return req, err
	}
	if err := transportHttp.DecodePostForm(r, req); err != nil {
		return req, err
	}

	if !transportHttp.IsFormRequest(r) {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil && err != io.EOF {
			return req, err
		}
	}
	return req, nil
}
func MakeAddTodoEndpoint(s TodolistService) transportHttp.Handler {
	return func(ctx context.Context, structReq interface{}) (interface{}, error) {
		req := structReq.(*dto.AddTodoReq)
		addTodoResp, err := s.AddTodo(ctx, req.AddTodoReq)
		var rs *dto.AddTodoResp
		if err != nil {
			rs = &dto.AddTodoResp{}
		} else {
			rs = &dto.AddTodoResp{AddTodoResp: addTodoResp}
		}
		return rs, err
	}
}

func DecodeAddTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	req := new(dto.AddTodoReq)
	// Order will be: URL parameters < query params < form data < JSON body
	// DEPRECATED: it is recommended to set 'http-parameter-priority: url' in grabkit.yaml
	if err := transportHttp.DecodeMuxVars(r, req); err != nil {
		return req, err
	}
	if err := transportHttp.DecodePostForm(r, req); err != nil {
		return req, err
	}

	if !transportHttp.IsFormRequest(r) {
		if err := json.NewDecoder(r.Body).Decode(req); err != nil && err != io.EOF {
			return req, err
		}
	}
	return req, nil
}
