package main

import (
	"github.com/lmfuture-ma/lmaker/cmd"
	"github.com/lmfuture-ma/lmaker/pkg/log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.RowMsg("main err", err)
	}
}
