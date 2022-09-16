package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var channel = make(chan int)

func producer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println("Produced: " + strconv.Itoa(i))
		channel <- i
	}
}

func consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		val := <-channel
		fmt.Println("Consumed: " + strconv.Itoa(val))
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(&wg)

	wg.Add(1)
	go consumer(&wg)

	wg.Wait()
	fmt.Println("Done")
}
