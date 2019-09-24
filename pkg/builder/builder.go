package builder

import (
	"bytes"
	"fmt"
	"github.com/lmfuture-ma/lmaker/pkg/templates"
	"io/ioutil"
	"os"
	"path"
)

type Builder interface {
	Build(td BuildData, templatePath, outPath string)
}

type BuilderImpl struct {
	templatesPath string
	outputPath    string
}

func NewBuilder(templatePath, outputPath string) BuilderImpl {
	return BuilderImpl{templatePath, outputPath}
}

func (bi *BuilderImpl) Build(td BuildData, templateFile, outFile string) error {
	var err error
	// get parse result in buf
	buf := new(bytes.Buffer)
	if err = templates.ParseTemplate(buf, path.Join(bi.templatesPath, templateFile), td); err != nil {
		return err
	}
	// buf to file
	outfilePath := path.Join(bi.outputPath, outFile)
	if _, err = os.Create(outfilePath); err != nil {
		return err
	}
	if err = ioutil.WriteFile(outfilePath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("write %s , err %v", outFile, err)
	}
	return nil
}
