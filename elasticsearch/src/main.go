package main // import "github.com/mapyo/GolangSample/elasticsearch"

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
)

func main() {
	app := cli.NewApp()

	app.Name = "elasticsearch sample"
	app.Usage = "elasticsearch sample command"
	app.Version = "0.0.1"
	app.Commands = commands
	app.Run(os.Args)
}

var commands = []cli.Command{
	createIndexAndMappingCommand,
}

var createIndexAndMappingCommand = cli.Command{
	Name:   "create",
	Usage:  "create index and mapping",
	Action: createIndexAndMapping,
}

func createIndexAndMapping(c *cli.Context) {
	fmt.Println("hoge")
}
