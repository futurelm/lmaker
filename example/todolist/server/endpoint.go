package server

import (
	"context"
	"fmt"

	"github.com/lmfuture-ma/lmaker/example/todolist/dto"
)

/*



// ListTodos...
func  listtodos(ctx context.Context) ([]*dto.Todo, err error) {
       return nil,nil
}
// GetTodo...
func  gettodo(ctx context.Context, int int64) (*dto.Todo, err error) {
       return nil,nil
}
// AddTodo...
func  addtodo(ctx context.Context, addTodoReq *dto.Todo) (*dto.Todo, err error) {
       return nil,nil
}
*/

var ts = []*dto.Todo{}

// ListTodos...
func listtodos(ctx context.Context) ([]*dto.Todo, error) {
	return ts, nil
}

// GetTodo...
func gettodo(ctx context.Context, int int64) (*dto.Todo, error) {
	for _, t := range ts {
		if t.Id == int {
			return t, nil
		}
	}
	return nil, fmt.Errorf("404")
}

func addtodo(ctx context.Context, todo *dto.Todo) (*dto.Todo, error) {
	ts = append(ts, todo)
	return todo, nil
}
