package main

import (
	"time"
	"log"
)

func heavyFunc(s string, second time.Duration, ch chan string) {
	time.Sleep(second * time.Second)
	log.Printf("hello %s", s)
	ch <- "finish"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go heavyFunc("world", 3, ch1)
	go heavyFunc("world2", 3, ch2)

	<-ch1
	log.Println("hoge")
	<-ch2
}