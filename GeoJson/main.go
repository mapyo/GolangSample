package main // import "github.com/mapyo/GolangSample/GeoJson"

import (
	"github.com/paulmach/go.geo"
	"fmt"
)

func main() {
	lineStringSample()
}

func lineStringSample() {
	p1 := &geo.Point{139.760820, 35.678340}
	p2 := &geo.Point{139.759750, 35.675270}

	geoJson := geo.NewLine(p1, p2).ToGeoJSON()

	encodedJSON, _ := geoJson.Geometry.MarshalJSON()

	fmt.Println(string(encodedJSON))
}
