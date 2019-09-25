package subCommand

import (
	"fmt"
	"github.com/lmfuture-ma/lmaker/model"
	"github.com/lmfuture-ma/lmaker/pkg"
	"github.com/lmfuture-ma/lmaker/pkg/log"
	"github.com/spf13/cobra"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

type DtoCommand BaseCommand

var dtoCommand *DtoCommand

func GetDto() *cobra.Command {
	var dto = &cobra.Command{
		Use:   "dto",
		Short: "dto-short",
		Long:  `dto-long`,
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

func delQuestion(c *DtoCommand, key string) kingpin.Action {
	return func(context *kingpin.ParseContext) error {
		delete(c.Question, key)
		return nil
	}
}

func (dc *DtoCommand) Run() error {
	//获取proto file
	curPath, _ := os.Getwd()
	files, err := pkg.FindProtoFile(curPath)
	if err != nil {
		log.RowMsg(err)
		return err
	}

	//解析proto
	//todo includePath variable
	protoObj, err := pkg.ParseProto(files[0], pkg.IncludePath)
	if err != nil {
		log.RowMsg(err)
		return err
	}
	//渲染dto模板
	fmt.Println(protoObj)
	return nil
}

func init() {
	dtoCommand = &DtoCommand{Config: model.LmakerInput{}, Question: qs}
}
