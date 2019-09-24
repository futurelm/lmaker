package protoc

import (
	"fmt"
	"strings"
)

func NewHttpRoute(p, method string) HTTPRoute {
	if !strings.HasPrefix(p, "/") {
		p = fmt.Sprintf("/%s", p)
	}
	return HTTPRoute{p, method}
}
