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
	sleepSecond := rand.Intn(3)
	time.Sleep(time.Duration(sleepSecond) * time.Second)
	fmt.Println("... ", id, " snore ...", "sleep", sleepSecond)
	wg.Done()
}
