package subCommand

import (
	"github.com/lmfuture-ma/lmaker/model"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"path/filepath"
)

type BaseCommand struct {
	Config model.LmakerInput
	// todo	Question map[string]map[string]*survey.Question
	Question map[string]*survey.Question
}

var baseTemplatePath string

var examplePB string

var qs = map[string]*survey.Question{
	"JSONOmitEmpty": {
		Name: "JSONOmitEmpty",
		Prompt: &survey.Confirm{
			Message: "Use omitempty tags in DTOs?",
			Default: true,
		},
	}, "DtoSqlTag": {
		Name: "DtoSqlTag",
		Prompt: &survey.Input{
			Default: "mysql",
			Message: "which db tag do u want to use?",
		},
	},
	"hello": {
		Name: "Hello",
		Prompt: &survey.Input{
			Message: "name a  name for hello",
		},
	},
}

func init() {
	examplePB = filepath.Join(os.Getenv("GOPATH"), "src", "github.com/lmfuture-ma/lmaker/pkg/templates/template_file/pb/todolist.proto")
	baseTemplatePath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com/lmfuture-ma/lmaker/pkg/templates/template_file")
}
