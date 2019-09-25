package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lmfuture-ma/lmaker/transprot"
	"net/http"

	"github.com/lmfuture-ma/lmaker/example/todolist/dto"
)

type Handler struct {
	ListTodos transprot.HttpHandler
	GetTodo   transprot.HttpHandler
	AddTodo   transprot.HttpHandler
}

func MakeHandler(service TodolistService, e *gin.Engine) *gin.Engine {
	handlers := registerHandler(service)
	e.Handle(http.MethodGet, "/listTodos", transprot.NewHttpServer(DecodeListTodosReq, handlers.ListTodos).Server)
	e.Handle(http.MethodGet, "/getTodo", transprot.NewHttpServer(DecodeGetTodoReq, handlers.GetTodo).Server)
	e.Handle(http.MethodPost, "/addTodo", transprot.NewHttpServer(DecodeAddTodoReq, handlers.AddTodo).Server)
	return e
}

func registerHandler(s TodolistService) Handler {
	var handler Handler
	handler.ListTodos = MakeListTodosEndpoint(s)
	handler.GetTodo = MakeGetTodoEndpoint(s)
	handler.AddTodo = MakeAddTodoEndpoint(s)
	return handler
}
func MakeListTodosEndpoint(s TodolistService) transprot.HttpHandler {
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

func DecodeListTodosReq(c *gin.Context) (interface{}, error) {
	return nil, nil

}
func MakeGetTodoEndpoint(s TodolistService) transprot.HttpHandler {
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

func DecodeGetTodoReq(c *gin.Context) (interface{}, error) {
	req := new(dto.GetTodoReq)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil

}
func MakeAddTodoEndpoint(s TodolistService) transprot.HttpHandler {
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

func DecodeAddTodoReq(c *gin.Context) (interface{}, error) {
	req := new(dto.AddTodoReq)
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, nil

}
