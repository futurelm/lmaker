package subCommand

import (
	"fmt"
	"github.com/go-openapi/inflect"
	"github.com/lmfuture-ma/lmaker/model"
	"github.com/lmfuture-ma/lmaker/pkg/builder"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"github.com/lmfuture-ma/lmaker/pkg/utils"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type createCMD struct {
	BaseCommand
	// template file path
	templatePath string
	// for go import  base path
	basePackage string
	// project path
	filePackage string
}

var createCommand *createCMD

func GetCreate() *cobra.Command {
	var create = &cobra.Command{
		Use:   "create",
		Short: "create-short",
		Long:  `create-long`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			printLog()
			if err := createCommand.Run(); err != nil {
				log.RowMsg(err)
			}
		},
	}
	create.PersistentFlags().StringVar(&createCommand.Config.ProjectName, "name", "example", "project name")
	return create
}

func (cc *createCMD) Run() error {
	if cc.Question != nil {
		if err := cc.Survey(); err != nil {
			return err
		}
	}
	if len(cc.Config.ProjectName) == 0 {
		return fmt.Errorf("projectName is nil")
	}
	fmt.Print("==========")
	formatName(cc)
	errArr := []error{}
	makeBaseProject(cc, &errArr)
	fmt.Print("==========")

	makeProtoFile(cc, &errArr)
	fmt.Print("==========")

	if len(errArr) > 0 {
		//删除当前目录
		log.RowMsg(fmt.Sprintf("%v，删除目录", errArr))
		err := os.RemoveAll(cc.Config.ProjectName)
		if err != nil {
			log.RowMsg("请手动删除")
		}
	}
	fmt.Println("==========(100%)")
	return nil
}

func makeProtoFile(cc *createCMD, err *[]error) {
	chdir(filepath.Join(cc.filePackage, cc.Config.ProjectName))
	utils.HandleErr(genCMD.Run(), err)
}

func formatName(cc *createCMD) {
	cc.Config.ProjectName = inflect.Underscore(cc.Config.ProjectName)
	cc.Config.ServiceName = inflect.Camelize(cc.Config.ProjectName)
}

func makeBaseProject(cc *createCMD, errArr *[]error) {
	cc.mkDir(errArr)
	if len(*errArr) > 0 {
		log.RowMsg(*errArr)
		return
	}
	cc.mkFile(errArr)
	if len(*errArr) > 0 {
		log.RowMsg(*errArr)
		return
	}
}

func (cc *createCMD) Survey() error {
	var qsArr []*survey.Question
	qsRequired := []string{"aaa"}
	for _, q := range qsRequired {
		if ques, ok := cc.BaseCommand.Question[q]; ok {
			qsArr = append(qsArr, ques)
		}
	}
	if err := survey.Ask(qsArr, &cc.Config); err != nil {
		log.RowMsg(err)
	}
	return nil
}

func (cc *createCMD) mkDir(flag *[]error) {
	utils.HandleErr(os.Mkdir(cc.Config.ProjectName, 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "cmd"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "cmd", "app"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "cmd", "conf"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "dto"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "server"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "server", "config"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "services"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "common"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "external"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "pb"), 0755), flag)
	utils.HandleErr(os.Mkdir(filepath.Join(cc.Config.ProjectName, "pkg"), 0755), flag)
}

func (cc *createCMD) mkFile(flag *[]error) {
	//cp example.proto
	utils.HandleErr(copyExampleProto(examplePB, filepath.Join(cc.filePackage, cc.Config.ProjectName, fmt.Sprintf("pb/%s.proto", cc.Config.ProjectName))), flag)
	bd := cc.baseBuildData()
	var build = builder.NewBuilder(cc.templatePath, filepath.Join(cc.filePackage, cc.Config.ProjectName))
	utils.HandleErr(build.Build(&bd, "main.tmpl", fmt.Sprintf("%s.go", cc.Config.ProjectName)), flag)
	utils.HandleErr(build.Build(&bd, "/app/server.tmpl", "/cmd/server.go"), flag)

}

func copyExampleProto(src string, dest string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = source.Close(), destination.Close()
	}()
	_, err = io.Copy(destination, source)
	return err
}

func (cc *createCMD) baseBuildData() builder.BuildDataImpl {
	var bd = builder.BuildDataImpl{ProjectName: cc.Config.ProjectName, ServiceName: cc.Config.ServiceName}
	bd.SetExtras([]string{cc.basePackage})
	return bd
}

func init() {
	pwd := utils.GetPwd()
	createCommand = &createCMD{
		BaseCommand:  BaseCommand{Question: qs, Config: model.LmakerInput{}},
		templatePath: baseTemplatePath,
		basePackage:  strings.TrimPrefix(pwd, fmt.Sprintf("%s/%s/", os.Getenv("GOPATH"), "src")),
		filePackage:  filepath.Join(pwd),
	}
}
