package command

import (
	"github.com/urfave/cli"
	"fmt"
)

var deleteIndexCommand = cli.Command{
	Name:   "delete",
	Usage:  "delete index",
	Action: deleteIndex,
}

func deleteIndex(c *cli.Context) {
	fmt.Println("deleteIndex!!!")
}

