package protoc

import (
	"fmt"
	"strings"
)

type File struct {
	Import  []string
	Name    string
	Service []Service
}

type Service struct {
	Name   string
	Method []Method
}

type Method struct {
	Name       string
	Params     []MethodField
	Results    []MethodField
	HTTPRoutes []HTTPRoute
	InputType  string
	OutputType string
}

type MethodField struct {
	Name string
	Type string
}

type HTTPRoute struct {
	Path       string
	HTTPMethod string
}

func NewHttpRoute(p, method string) HTTPRoute {
	if !strings.HasPrefix(p, "/") {
		p = fmt.Sprintf("/%s", p)
	}
	return HTTPRoute{p, method}
}
