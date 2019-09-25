package protoc

import (
	"fmt"
	"strings"
)

type File struct {
	Import  []string  `json:"import"`
	Name    string    `json:"name"`
	Service []Service `json:"services"`
	Message []Message `json:"messages"`
}

type Service struct {
	Name   string   `json:"name"`
	Method []Method `json:"methods"`
}

type Method struct {
	Name       string      `json:"name"`
	Params     []Field     `json:"params"`
	Results    []Field     `json:"result"`
	HTTPRoutes []HTTPRoute `json:"http_routes"`
	InputType  string      `json:"input_type"`
	OutputType string      `json:"out_type"`
}

type Message struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	JsonTag string `json:"json_tag"`
}

type HTTPRoute struct {
	Path       string `json:"path"`
	HTTPMethod string `json:"http_method"`
}

func NewHttpRoute(p, method string) HTTPRoute {
	if !strings.HasPrefix(p, "/") {
		p = fmt.Sprintf("/%s", p)
	}
	return HTTPRoute{p, method}
}
