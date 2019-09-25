package services

import (
	"context"
	"github.com/lmfuture-ma/lmaker/example/todolist/dto"
)

type TodolistService interface {
	// ListTodos...
	ListTodos(ctx context.Context) (todos []*dto.Todo, err error)
	// GetTodo...
	GetTodo(ctx context.Context, int int64) (todo *dto.Todo, err error)
	// AddTodo...
	AddTodo(ctx context.Context, addTodoReq *dto.Todo) (addTodoResp *dto.Todo, err error)
}
