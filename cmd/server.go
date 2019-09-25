package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/lmfuture-ma/lmaker/cmd/subCommand"
)

var cobraCMD *cobra.Command

func Run() error {
	return cobraCMD.Execute()
}

func init() {
	cobraCMD = &cobra.Command{
		Use:     "lmaker",
		Version: "1111",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("cobracmd")
		},
	}
	cobraCMD.AddCommand(subCommand.GetDto())
	cobraCMD.AddCommand(subCommand.GetCreate())
	cobraCMD.AddCommand(subCommand.GetGen())
}
