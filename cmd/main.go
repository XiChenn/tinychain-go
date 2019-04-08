package main

import (
	"tinychian-go/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.GetDb().Close()

	cli := CLI{bc}
	cli.Run()
}
