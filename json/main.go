package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	decodeSample()

	encodeSample()

}

func decodeSample() {
	bytes, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	// decode
	var parsons []Person
	json.Unmarshal(bytes, &parsons)

	for _, p := range parsons {
		fmt.Printf("%d: %s\n", p.Id, p.Name)
	}
}

func encodeSample() {
	// encodeSample
	parsonsSample := []Person{
		{
			Id:   3,
			Name: "foofoo",
		},
		{
			Id:   4,
			Name: "hahaha",
		},
	}

	bytes, err := json.Marshal(parsonsSample)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
