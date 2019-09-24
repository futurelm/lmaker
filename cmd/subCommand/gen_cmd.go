package subCommand

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/lmfuture-ma/lmaker/model"
	"github.com/lmfuture-ma/lmaker/pkg"
	"github.com/lmfuture-ma/lmaker/pkg/builder"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"github.com/lmfuture-ma/lmaker/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

type genCommand struct {
	BaseCommand
	// template file path
	templatePath string
	// for go import  base path
	basePackage string
	// project path
	filePackage string
	// workdir pb file
	pbPath string
	//
	fromCreate bool
}

var genCMD *genCommand

func GetGen() *cobra.Command {
	var dto = &cobra.Command{
		Use:   "gen",
		Short: "gen",
		Long:  `gen-long`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			if err := dtoCommand.Run(); err != nil {
				log.RowMsg(err)
			}
		},
	}
	dto.PersistentFlags().StringVar(&dtoCommand.Config.DtoSqlTag, "tag", "", "json tag")
	return dto
}

func (g *genCommand) Run() error {
	files, err := pkg.FindProtoFile(g.filePackage)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("proto file not found")
	}
	err = g.protoc(files[0])
	if err != nil {
		return err
	}
	return nil
}

func (g *genCommand) protoc(path string) error {
	//pb.go
	err := pkg.ParseProtoToFile(path, pkg.IncludePath)
	if err != nil {
		return err
	}
	descriptorSet, err := pkg.ParseProto(path, pkg.IncludePath)
	if err != nil {
		return err
	}
	errArr := []error{}
	g.generateFile(descriptorSet, &errArr)
	if len(errArr) > 0 {
		return fmt.Errorf("err %v", errArr)
	}
	return nil
}

func (g *genCommand) generateFile(descriptorSet *descriptor.FileDescriptorSet, errArr *[]error) {
	var build = builder.NewBuilder(g.templatePath, g.filePackage)
	bd := g.baseBuildData()
	err := bd.ReadProtoFile(*descriptorSet.File[0])
	if err != nil {
		*errArr = append(*errArr, err)
		return
	}
	log.RowMsg(fmt.Sprintf("%+v", bd))
	utils.HandleErr(build.Build(&bd, "/services/handler.tmpl", "/services/z_handler.go"), errArr)
	utils.HandleErr(build.Build(&bd, "/services/services.tmpl", "/services/z_service.go"), errArr)
	if g.fromCreate {
		utils.HandleErr(build.Build(&bd, "/server/server.tmpl", "/server/z_server.go"), errArr)
		utils.HandleErr(build.Build(&bd, "/server/endpoint.tmpl", "/server/endpoint.go"), errArr)

	}

}

func (cc *genCommand) baseBuildData() builder.BuildDataImpl {
	var bd = builder.BuildDataImpl{ProjectName: cc.Config.ProjectName, ServiceName: cc.Config.ServiceName}
	bd.FilePackage = cc.filePackage
	bd.ProjectName = strings.Split(cc.filePackage, "/")[len(strings.Split(cc.filePackage, "/"))-1]
	bd.SetExtras([]string{cc.basePackage})
	return bd
}

func init() {
	pwd := utils.GetPwd()
	genCMD = &genCommand{
		BaseCommand:  BaseCommand{Question: qs, Config: model.LmakerInput{}},
		templatePath: baseTemplatePath,
		basePackage:  strings.TrimPrefix(pwd, fmt.Sprintf("%s/%s/", os.Getenv("GOPATH"), "src")),
		filePackage:  filepath.Join(pwd),
	}
}

func chdir(s string) {
	genCMD.filePackage = s
	genCMD.basePackage = strings.TrimPrefix(s, fmt.Sprintf("%s/%s/", os.Getenv("GOPATH"), "src"))
	genCMD.fromCreate = true
}
