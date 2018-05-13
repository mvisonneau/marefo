package main

import (
	"os"

	"github.com/mvisonneau/marefo/cli"
)

var version = "<devel>"

func main() {
	cli.Run(&version).Run(os.Args)
}
