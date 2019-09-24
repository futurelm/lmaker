package protoc

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
