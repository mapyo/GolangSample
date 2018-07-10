package main // import "github.com/mapyo/GolangSample/elasticsearch"

import (
	"github.com/urfave/cli"
	"os"
	"github.com/mapyo/GolangSample/elasticsearch/src/command"
)

func main() {
	app := cli.NewApp()

	app.Name = "elasticsearch sample"
	app.Usage = "elasticsearch sample command"
	app.Version = "0.0.1"
	app.Commands = command.Commands
	app.Run(os.Args)
}
