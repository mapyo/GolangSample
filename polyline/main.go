package main

import (
	"fmt"
	"github.com/paulmach/go.geo"
)

func main() {
	// decode
	path := geo.NewPathFromEncoding("slwxEc``tYdRtEsM|f@dBvV}LnMgBlIsN`GyHr@cY_FaGqNL}LiBwM{DoLiCmMbAkW`BmH`GoLdYtE")

	fmt.Println(path)

	for _, v := range path.PointSet {
		fmt.Printf("%f,%f,0\n", v[0], v[1])
	}
}
