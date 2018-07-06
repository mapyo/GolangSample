package command

import "github.com/urfave/cli"

var Commands = []cli.Command{
	createIndexAndMappingCommand,
	deleteIndexCommand,
	bulkInsertCommand,
	searchCommand,
}

