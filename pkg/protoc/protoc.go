package protoc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var IncludePath []string

func ParseProto(protoFiles string, includePath []string) (*descriptor.FileDescriptorSet, error) {
	var rs = &descriptor.FileDescriptorSet{}
	pwd, err := os.Getwd()
	if err != nil {
		log.RowMsg(err)
		return nil, err
	}
	tmpFile, err := ioutil.TempFile(pwd, "*.txt")
	if err != nil {
		return &descriptor.FileDescriptorSet{}, err
	}
	defer func() {
		if err := tmpFile.Close(); err != nil {
			log.RowMsg(err)
		}
		if err:=os.Remove(tmpFile.Name()); err != nil {
			log.RowMsg(err)
		}
	}()
	//构造protoc 命令
	args := []string{}
	args = append(args, "-o", tmpFile.Name()) //output file
	// /Users/liangliang.ma/gopath/src/github.com/lmfuture-ma/golearning/lmaker/pb/prd.proto -> /Users/liangliang.ma/gopath/src/github.com/lmfuture-ma/golearning/lmaker/pb
	protoPath := string([]byte(protoFiles)[:strings.LastIndex(protoFiles, "/")])
	args = append(args, "--proto_path", protoPath) //output file
	args = append(args, "--go_out", protoPath)
	//todo include common path
	for _, p := range includePath {
		args = append(args, "-I", p)
	}
	args = append(args, protoFiles)
	c := exec.Command("protoc", args...)
	if out, err := c.CombinedOutput(); err != nil {
		log.RowMsg(fmt.Sprintf("Failed to run command: %s\n\nError: %v\n\nOutput: %s\n", strings.Join(c.Args, " "), err, out))
	}
	bytes, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}
	if err = proto.Unmarshal(bytes, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func ParseProtoToFile(protoFiles string, includePath []string) error {
	c := exec.Command("protoc", generateCommand(protoFiles, includePath)...)
	if out, err := c.CombinedOutput(); err != nil {
		log.RowMsg(fmt.Sprintf("Failed to run command: %s\n\nError: %v\n\nOutput: %s\n", strings.Join(c.Args, " "), err, out))
		return err
	}
	return nil
}
func generateCommand(protoFiles string, includePath []string) []string {
	args := []string{}
	protoPath := string([]byte(protoFiles)[:strings.LastIndex(protoFiles, "/")])
	args = append(args, "--proto_path", protoPath) //output file
	args = append(args, "--go_out", protoPath)
	//todo include common path
	for _, p := range includePath {
		args = append(args, "-I", p)
	}
	args = append(args, protoFiles)
	return args
}

func init() {
	IncludePath = []string{filepath.Join(os.Getenv("GOPATH"), "src", "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis")}
}
