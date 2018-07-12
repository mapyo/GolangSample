package command

import (
	"github.com/urfave/cli"
	"github.com/olivere/elastic"
	"context"
	"fmt"
)

var createIndexAndMappingCommand = cli.Command{
	Name:   "create",
	Usage:  "create index and mapping",
	Action: createIndexAndMapping,
}

func createIndexAndMapping(c *cli.Context) {
	client, err := elastic.NewSimpleClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	defer client.Stop()

	mapping := `
{
 "settings":{
   "number_of_shards":1,
   "number_of_replicas":0
 },
 "mappings": {
   "type": {
     "properties": {
       "name": {
         "type": "keyword"
       },
       "location": {
         "type": "geo_point"
       }
     }
   }
 }
}
`

	ctx := context.Background()
	createIndex, err := client.CreateIndex("mapyo").BodyString(mapping).Do(ctx)
	fmt.Println(createIndex)
	if err != nil {
		// Handle error
		panic(err)
	}

	if !createIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	} else {
		fmt.Println("success")
	}
}
