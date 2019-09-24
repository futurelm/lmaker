package templates

import (
	"bytes"
	"fmt"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"path/filepath"
	"strings"
	"text/template"
)

func ParseTemplate(buff *bytes.Buffer, templatePath string, data interface{}) error {
	tmpl := getTemplate(templatePath)
	if tmpl == nil {
		return fmt.Errorf("parse file error")
	}
	err := tmpl.Execute(buff, data)
	if err != nil {
		log.RowMsg("tmpl Execute  err ")
		return err
	}
	return nil
}

func getTemplate(templatePath string) *template.Template {
	//tmpl,err:=template.New("parent.tmpl").Funcs(getFuncMap()).ParseFiles("/Users/liangliang.ma/gopath/src/github.com/lmfuture-ma/golearning/lmaker/pkg/templates/template_file/parent.tmpl")
	//if err != nil {
	//	log.RowMsg(err)
	//	return nil
	//}
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(getFuncMap()).ParseFiles(templatePath)
	if err != nil {
		log.RowMsg(err)
		return nil
	}
	return tmpl
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"firstUpperCase": strings.Title,
		// don't firstLower, temp use ToLower
		"loweCase": strings.ToLower,
	}
}
