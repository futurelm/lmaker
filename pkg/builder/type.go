package builder

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"github.com/lmfuture-ma/lmaker/pkg/protoc"
	"google.golang.org/genproto/googleapis/api/annotations"
	"strings"
)

type BuildData interface {
	SetExtras([]string)
	Dup() BuildData
}

type BuildDataImpl struct {
	ProjectName string
	ServiceName string
	Extra       []string
	File        protoc.File
	FilePackage string
}

func (bd *BuildDataImpl) SetExtras(ext []string) {
	if len(ext) > 0 {
		bd.Extra = ext
	}
}

func (bd *BuildDataImpl) Dup() BuildData {
	return bd
}

func (bd *BuildDataImpl) ReadProtoFile(file descriptor.FileDescriptorProto) error {
	var f = protoc.File{}
	f.Import = file.Dependency
	f.Name = file.GetName()
	//meg map
	var msgMap = map[string]*descriptor.DescriptorProto{}
	for _, msg := range file.MessageType {
		if _, ok := msgMap[msg.GetName()]; !ok {
			msgMap[msg.GetName()] = msg
		}
	}
	for _, svc := range file.Service {
		var protoSVC = protoc.Service{Name: svc.GetName()}
		for _, method := range svc.Method {
			var m = protoc.Method{}
			m.Name = method.GetName()
			m.HTTPRoutes = getHttpRoute(method)
			m.InputType = formatInput(method.GetInputType())
			m.OutputType = formatInput(method.GetOutputType())
			m.Params = readFiled(method.GetInputType(), msgMap)
			m.Results = readFiled(method.GetOutputType(), msgMap)
			protoSVC.Method = append(protoSVC.Method, m)
		}
		f.Service = append(f.Service, protoSVC)
	}
	bd.File = f
	return nil
}

// work formatOutput too
func formatInput(s string) string {
	//input_type:".todolistpb.GetTodoReq"
	//input_type:".google.protobuf.Empty"
	oneWord := getTypeFromProtoTypeName(s)
	if oneWord == "Empty" {
		return ""
	}
	return fmt.Sprintf("pb.%s", oneWord)
}

func readFiled(s string, msgMap map[string]*descriptor.DescriptorProto) []protoc.MethodField {
	var rs = []protoc.MethodField{}
	// s = ".todolistpb.GetTodoReq"
	var desc *descriptor.DescriptorProto
	desc, ok := msgMap[getTypeFromProtoTypeName(s)]
	if !ok {
		return rs
	}
	for _, filed := range desc.GetField() {
		rs = append(rs, protoc.MethodField{filed.GetName(), goType(filed)})
	}
	return rs
}

func getHttpRoute(method *descriptor.MethodDescriptorProto) []protoc.HTTPRoute {
	var httpRoutes = []protoc.HTTPRoute{}
	opts := method.GetOptions()
	extension, err := proto.GetExtension(opts, annotations.E_Http)
	if err != nil {
		return httpRoutes
	}
	if rule, ok := extension.(*annotations.HttpRule); ok {
		if p := rule.GetPost(); p != "" {
			httpRoutes = append(httpRoutes, protoc.NewHttpRoute(p, "http.MethodPost"))
		}
		if p := rule.GetGet(); p != "" {
			httpRoutes = append(httpRoutes, protoc.NewHttpRoute(p, "http.MethodGet"))
		}
		if p := rule.GetDelete(); p != "" {
			httpRoutes = append(httpRoutes, protoc.NewHttpRoute(p, "http.MethodDelete"))
		}
		if p := rule.GetPatch(); p != "" {
			httpRoutes = append(httpRoutes, protoc.NewHttpRoute(p, "http.MethodPatch"))
		}
		if p := rule.GetPut(); p != "" {
			httpRoutes = append(httpRoutes, protoc.NewHttpRoute(p, "http.MethodPut"))
		}
	}
	return httpRoutes
}

func getTypeFromProtoTypeName(s string) string {
	ar := strings.Split(s, ".")
	if len(ar) == 0 {
		return ""
	}
	return ar[len(strings.Split(s, "."))-1]
}

// todo move these method to ./protoc

// goType returns a string representing the type name
func goType(field *descriptor.FieldDescriptorProto) (typ string) {

	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		typ = "float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		typ = "float32"
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		typ = "int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		typ = "uint64"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		typ = "int32"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		typ = "uint32"
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		typ = "uint64"
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		typ = "uint32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		typ = "bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		typ = "string"
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// unsupported (groups are deprecated)
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		switch field.GetTypeName() {
		case ".google.protobuf.Timestamp":
			typ = "*time.Time"
		case ".google.protobuf.Any":
			typ = "*any.Any"
		default:
			typ = fmt.Sprintf("*pb.%s", getTypeFromProtoTypeName(field.GetTypeName()))
		}
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		typ = "unknown"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		typ = "[]byte"
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		typ = "int32"
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		typ = "int64"
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		typ = "int32"
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		typ = "int64"
	default:
	}

	if isRepeated(field) {
		typ = "[]" + typ
	}
	log.RowMsg(fmt.Sprintf("before %v type %s", field, typ))

	return
}

func isRepeated(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED
}
