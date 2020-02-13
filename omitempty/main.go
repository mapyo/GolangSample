package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	Bar []string `json:",omitempty"`
}

func main() {
	b := []byte(`{"bar": ["a", "b", "c"]}`)
	var f foo
	if err := json.Unmarshal(b, &f); err != nil {
		panic(err)
	}
	fmt.Println(f)

	f2 := foo{Bar: []string{}}
	b2, err := json.Marshal(f2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b2))
}
