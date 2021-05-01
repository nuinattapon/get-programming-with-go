package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var finishedCounter int64 = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 5)
	for i := 0; i < 20; i++ {
		go sleepyGopher(i, ch)
	}

	timeOutSecond := 3
	timeout := time.After(time.Duration(timeOutSecond) * time.Second)
	for i := 0; i < 20; i++ {
		select {
		case gopherID := <-ch:
			fmt.Println("gopher", gopherID, "has finished sleeping")
			atomic.AddInt64(&finishedCounter, 1)
		case <-timeout:
			fmt.Printf("My patience ran out within %d seconds\n", timeOutSecond)
			fmt.Printf("# of finished gopher is %d \n", finishedCounter)
			return
		}
	}
	// waitGroup.Wait()
}
func sleepyGopher(id int, ch chan<- int) {
	sleepDuration := rand.Intn(4000)
	fmt.Printf("...%2d snoor... sleep %4d\n", id, sleepDuration)

	time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	ch <- id
}
