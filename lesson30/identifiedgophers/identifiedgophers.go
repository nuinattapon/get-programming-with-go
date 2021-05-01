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

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sleepyGopher(i)
	}
	// time.Sleep(4 * time.Second)
	wg.Wait()
}

func sleepyGopher(id int) {
	sleepSecond := rand.Intn(4)
	fmt.Println("... ", id, "snore ...", "sleep", sleepSecond)
	time.Sleep(time.Duration(sleepSecond) * time.Second)
	wg.Done()
}
