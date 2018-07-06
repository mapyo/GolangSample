package command

import (
	"github.com/urfave/cli"
	"fmt"
)

var bulkInsertCommand = cli.Command{
	Name:   "bulk",
	Usage:  "bulk insert something data",
	Action: bulkInsert,
}

func bulkInsert(c *cli.Context) {
	fmt.Println("bulkInsert!!!")
}

