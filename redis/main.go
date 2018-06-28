package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

const GeoKey = "geo_sample"

var locations = []redis.GeoLocation{
	{
		Name:      "Tokyo Tower",
		Latitude:  35.658582,
		Longitude: 139.745464,
	},
	{
		Name:      "SkyTree",
		Latitude:  35.710033,
		Longitude: 139.810716,
	},
	{
		Name:      "Yokohama Marine Tower",
		Latitude:  35.443938,
		Longitude: 139.650941,
	},
}

func main() {
	// ref. https://github.com/go-redis/redis#quickstart
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	for _, geoLocation := range locations {
		fmt.Println(geoLocation)
		res, err := client.GeoAdd(GeoKey, &geoLocation).Result()
		fmt.Println(res, err)
	}

	query := redis.GeoRadiusQuery{Radius: 10, Unit: "km", Sort: "ASC"}

	res, err := client.GeoRadius(GeoKey, 139.766247, 35.68129, &query).Result()
	fmt.Println(res, err)
	// Output: [{Tokyo Tower 0 0 0 0} {SkyTree 0 0 0 0}] <nil>
}
