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

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sleepyGopher(i, ch)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-ch
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
	wg.Wait()
}

func sleepyGopher(id int, ch chan<- int) {
	sleepDuration := rand.Intn(4000)
	fmt.Println("... ", id, "snore ...", "sleep", sleepDuration)

	time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	// fmt.Println("... ", id, " snore ...")
	ch <- id
	wg.Done()

}
