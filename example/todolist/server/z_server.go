package server

import (
	"context"
	"github.com/lmfuture-ma/lmaker/example/todolist/dto"
	"github.com/lmfuture-ma/lmaker/example/todolist/services"
)

func NewServiceImpl() services.TodolistService {
	return &todolistServiceImpl{}
}

type todolistServiceImpl struct{}

// ListTodos...
func (s *todolistServiceImpl) ListTodos(ctx context.Context) ([]*dto.Todo, error) {
	return listtodos(ctx)
}

// GetTodo...
func (s *todolistServiceImpl) GetTodo(ctx context.Context, int int64) (*dto.Todo, error) {
	return gettodo(ctx, int)
}

// AddTodo...
func (s *todolistServiceImpl) AddTodo(ctx context.Context, addTodoReq *dto.Todo) (*dto.Todo, error) {
	return addtodo(ctx, addTodoReq)
}
