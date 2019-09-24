package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/lmfuture-ma/lmaker/cmd/subCommand"
)

const log = `
 ___      __   __  _______  ___   _  _______  ______           _______  _______ 
|   |    |  |_|  ||   _   ||   | | ||       ||    _ |         |       ||       |
|   |    |       ||  |_|  ||   |_| ||    ___||   | ||   ____  |    ___||   _   |
|   |    |       ||       ||      _||   |___ |   |_||_ |____| |   | __ |  | |  |
|   |___ |       ||       ||     |_ |    ___||    __  |       |   ||  ||  |_|  |
|       || ||_|| ||   _   ||    _  ||   |___ |   |  | |       |   |_| ||       |
|_______||_|   |_||__| |__||___| |_||_______||___|  |_|       |_______||_______|
`

var cobraCMD *cobra.Command

func Run() error {
	return cobraCMD.Execute()
}

func print() {
	fmt.Println(color.BlueString(log))
	fmt.Println()
	fmt.Println("🍺-----------------------🍺---------------------------🍺----------------------🍺")
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
