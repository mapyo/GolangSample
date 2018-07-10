package command

import (
	"github.com/urfave/cli"
	"fmt"
)

var createIndexAndMappingCommand = cli.Command{
	Name:   "create",
	Usage:  "create index and mapping",
	Action: createIndexAndMapping,
}

func createIndexAndMapping(c *cli.Context) {
	fmt.Println("createIndexAndMapping!!!")
}
