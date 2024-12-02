package main

import (
	"fmt"
	"sync"
)

// Custome type to encapsulate the logic for concurrent safe data manipulation
type SafeCounter struct {
	count      int
	sync.Mutex //inherited
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.count += delta
	}
	sf.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()

	fmt.Println("count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
