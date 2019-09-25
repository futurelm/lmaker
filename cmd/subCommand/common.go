package subCommand

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lmfuture-ma/lmaker/model"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"path/filepath"
)

const logo = `
 ___      __   __  _______  ___   _  _______  ______           _______  _______ 
|   |    |  |_|  ||   _   ||   | | ||       ||    _ |         |       ||       |
|   |    |       ||  |_|  ||   |_| ||    ___||   | ||   ____  |    ___||   _   |
|   |    |       ||       ||      _||   |___ |   |_||_ |____| |   | __ |  | |  |
|   |___ |       ||       ||     |_ |    ___||    __  |       |   ||  ||  |_|  |
|       || ||_|| ||   _   ||    _  ||   |___ |   |  | |       |   |_| ||       |
|_______||_|   |_||__| |__||___| |_||_______||___|  |_|       |_______||_______|
`

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

func printLog() {
	fmt.Println(color.BlueString(logo))
	fmt.Println("🍺-----------------------🍺---------------------------🍺----------------------🍺")
}
