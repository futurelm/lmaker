package dto

type AddTodoResp struct {
	AddTodoResp *Todo `json:"addTodoResp"`
}
type AddTodoReq struct {
	AddTodoReq *Todo `json:"addTodoReq"`
}
type GetTodoResp struct {
	Todo *Todo `json:"todo"`
}
type Todo struct {
	Id          int64  `json:"id"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Created_at  string `json:"createdAt"`
}
type GetTodoReq struct {
	Int int64 `json:"int"`
}
type ListTodosResponse struct {
	Todos []*Todo `json:"todos"`
}
