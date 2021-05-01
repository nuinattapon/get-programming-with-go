package main

import (
	"fmt"
	"sync"
)

// Visited tracks whether web pages have been visited.
// Its methods may be used concurrently with one another.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

// VisitLink tracks that the page with the given URL has
// been visited, and returns the updated link count.
func (v *Visited) VisitLink(url string) int {
	defer v.mu.Unlock()

	v.mu.Lock()
	if v.visited == nil {
		v.visited = map[string]int{}
	}
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {
	// var visited Visited = Visited{mu: sync.Mutex{}, visited: map[string]int{}}
	visited := Visited{mu: sync.Mutex{}, visited: map[string]int{}}
	fmt.Printf("%+v\n", visited.visited)
	visited.VisitLink("https://gopl.io")
	fmt.Printf("%+v\n", visited.visited)
	visited.VisitLink("https://https://golang.org/")
	fmt.Printf("%+v\n", visited.visited)
	visited.VisitLink("https://gopl.io")
	fmt.Printf("%+v\n", visited.visited)
	visited.VisitLink("https://gopl.io")
	fmt.Printf("%+v\n", visited.visited)
	visited.VisitLink("https://https://golang.org/")
	fmt.Printf("%+v\n", visited.visited)
}
