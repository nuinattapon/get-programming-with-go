package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sleepyGopher(i, ch)
	}

	timeOutSecond := 3
	timeout := time.After(time.Duration(timeOutSecond) * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case gopherID := <-ch:
			fmt.Println("gopher", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Printf("my patience ran out within %d seconds\n", timeOutSecond)
			return
		}
	}
	wg.Wait()
}
func sleepyGopher(id int, ch chan<- int) {
	sleepDuration := rand.Intn(4000)
	fmt.Printf("...%2d snoor... sleep %4d\n", id, sleepDuration)

	time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	ch <- id
	wg.Done()
}
