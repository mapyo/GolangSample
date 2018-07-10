package command

import (
	"github.com/urfave/cli"
	"fmt"
)

var searchCommand = cli.Command{
	Name:   "search",
	Usage:  "search data",
	Action: search,
}

func search(c *cli.Context) {
	fmt.Println("search!!!")
}
