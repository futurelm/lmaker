package builder

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lmfuture-ma/lmaker/pkg/protoc"
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
	f.Name = file.GetName()
	//meg map
	var msgArr = []protoc.Message{}
	var msgMap = map[string]*descriptor.DescriptorProto{}
	for _, msg := range file.MessageType {
		if _, ok := msgMap[msg.GetName()]; !ok {
			msgMap[msg.GetName()] = msg
			msgArr = append(msgArr, protoc.ReadMessage(msg))
		}
	}
	f.Message = msgArr
	for _, svc := range file.Service {
		var protoSVC = protoc.Service{Name: svc.GetName()}
		for _, method := range svc.Method {
			var m = protoc.Method{}
			m.Name = method.GetName()
			m.HTTPRoutes = protoc.GetHttpRoute(method)
			m.InputType = protoc.FormatInput(method.GetInputType())
			m.OutputType = protoc.FormatInput(method.GetOutputType())
			m.Params = protoc.ReadMethodFiled(method.GetInputType(), msgMap)
			m.Results = protoc.ReadMethodFiled(method.GetOutputType(), msgMap)
			protoSVC.Method = append(protoSVC.Method, m)
		}
		f.Service = append(f.Service, protoSVC)
	}
	bd.File = f
	return nil
}
