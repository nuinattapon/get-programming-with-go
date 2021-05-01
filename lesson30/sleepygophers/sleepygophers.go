package main

import (
	"fmt"
	"time"
	"sync"

)

var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sleepyGopher()
	}
	wg.Wait()
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
	wg.Done()

}
