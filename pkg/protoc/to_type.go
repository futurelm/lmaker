package protoc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/genproto/googleapis/api/annotations"
	"strings"
)

// work formatOutput too
func FormatInput(s string) string {
	//input_type:".todolistpb.GetTodoReq"
	//input_type:".google.protobuf.Empty"
	oneWord := getTypeFromProtoTypeName(s)
	if oneWord == "Empty" {
		return ""
	}
	return fmt.Sprintf("dto.%s", oneWord)
}

func ReadMethodFiled(s string, msgMap map[string]*descriptor.DescriptorProto) []Field {
	var rs = []Field{}
	// s = ".todolistpb.GetTodoReq"
	var desc *descriptor.DescriptorProto
	desc, ok := msgMap[getTypeFromProtoTypeName(s)]
	if !ok {
		return rs
	}
	for _, filed := range desc.GetField() {
		rs = append(rs, Field{filed.GetName(), goType(filed, true), filed.GetJsonName()})
	}
	return rs
}

func ReadMessage(msg *descriptor.DescriptorProto) Message {
	var m = Message{Name: strings.Title(msg.GetName())}
	var fields = []Field{}
	for _, f := range msg.Field {
		fields = append(fields, Field{Name: strings.Title(f.GetName()), Type: goType(f, false), JsonTag: f.GetJsonName()})
	}
	m.Fields = fields
	return m
}

func GetHttpRoute(method *descriptor.MethodDescriptorProto) []HTTPRoute {
	var httpRoutes = []HTTPRoute{}
	opts := method.GetOptions()
	extension, err := proto.GetExtension(opts, annotations.E_Http)
	if err != nil {
		return httpRoutes
	}
	if rule, ok := extension.(*annotations.HttpRule); ok {
		if p := rule.GetPost(); p != "" {
			httpRoutes = append(httpRoutes, NewHttpRoute(p, "http.MethodPost"))
		}
		if p := rule.GetGet(); p != "" {
			httpRoutes = append(httpRoutes, NewHttpRoute(p, "http.MethodGet"))
		}
		if p := rule.GetDelete(); p != "" {
			httpRoutes = append(httpRoutes, NewHttpRoute(p, "http.MethodDelete"))
		}
		if p := rule.GetPatch(); p != "" {
			httpRoutes = append(httpRoutes, NewHttpRoute(p, "http.MethodPatch"))
		}
		if p := rule.GetPut(); p != "" {
			httpRoutes = append(httpRoutes, NewHttpRoute(p, "http.MethodPut"))
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
func goType(field *descriptor.FieldDescriptorProto, prefix bool) (typ string) {

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
			typ = fmt.Sprintf("*%s", getTypeFromProtoTypeName(field.GetTypeName()))
			if prefix {
				typ = fmt.Sprintf("*dto.%s", getTypeFromProtoTypeName(field.GetTypeName()))
			}
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
	return
}

func isRepeated(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED
}
