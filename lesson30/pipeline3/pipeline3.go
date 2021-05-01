package main

import (
	"fmt"
	"strings"
)

func main() {
	ch0 := make(chan string)
	ch1 := make(chan string)
	go sourceGopher(ch0)
	go filterGopher(ch0, ch1)
	printGopher(ch1)
}
func sourceGopher(downstream chan<- string) {
	for _, v := range []string{"hello world", "a bad apple", "hi there", "how r u?", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}
func filterGopher(upstream <-chan string, downstream chan<- string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}
func printGopher(upstream <-chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
