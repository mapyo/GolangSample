package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	size := 11
	hoges := make([]string, size)

	for i := 0; i < size; i++ {
		hoges[i] = fmt.Sprintf("hoge%d", i)
	}

	limit := make(chan int, 10)

	var wg sync.WaitGroup
	for _, v := range hoges {
		wg.Add(1)
		go func(v string) {
			limit <- 1

			defer wg.Done()
			doSomething(v)

			<-limit
		}(v)
	}

	wg.Wait()
}