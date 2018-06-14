package main

import (
	"github.com/paulmach/go.geo"
	"fmt"
)

func main() {
	decode()
	encode()
}

func encode() {
	path := geo.Path{
		PointSet: geo.PointSet{
			geo.Point{139.760820, 35.678340},
			geo.Point{139.759750, 35.675270},
		},
	}

	polyline := path.Encode()
	fmt.Println(polyline)
}

func decode() {
	path := geo.NewPathFromEncoding("slwxEc``tYdRtEsM|f@dBvV}LnMgBlIsN`GyHr@cY_FaGqNL}LiBwM{DoLiCmMbAkW`BmH`GoLdYtE")

	fmt.Println(path)

	for _, v := range path.PointSet {
		fmt.Printf("%f,%f,0\n", v[0], v[1])
	}
}
