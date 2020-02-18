package main

import (
	"fmt"
	"github.com/paulmach/go.geo"
	"github.com/paulmach/orb/geojson"
)

func main() {
	lineStringSample()
	parseGeoJSON()
}

func parseGeoJSON() {
	rawFeatureJSON := []byte(`
{
  "type": "Feature",
  "geometry": {
    "type": "Polygon",
    "coordinates": [
      [
        [ 100, 0 ],
        [ 101, 0 ],
        [ 101, 1 ],
        [ 100, 1 ],
        [ 100, 0 ]
      ]
    ]
  }
}`)
	fc1, err := geojson.UnmarshalFeature(rawFeatureJSON)
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed parse geojson")
	}
	fmt.Printf("%+v\n", fc1)
}

func lineStringSample() {
	p1 := &geo.Point{139.760820, 35.678340}
	p2 := &geo.Point{139.759750, 35.675270}

	geoJson := geo.NewLine(p1, p2).ToGeoJSON()

	encodedJSON, _ := geoJson.Geometry.MarshalJSON()

	fmt.Println(string(encodedJSON))
}
