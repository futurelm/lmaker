package protoc

import (
	"fmt"
	"os"
	"path"
	"strings"
)

//FindProtoFile 寻找path目录下pb目录里的proto文件
func FindProtoFile(curPath string) ([]string, error) {
	rs := []string{}
	filePath := path.Join(curPath, "pb")
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("can't find %s dir", filePath)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fs, err := f.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	for _, f := range fs {
		if strings.HasSuffix(f, ".proto") {
			rs = append(rs, path.Join(filePath, f))
		}
	}
	return rs, nil
}
