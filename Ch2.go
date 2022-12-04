package main

import (
	"fmt"
	"sync"
)

func say(wg *sync.WaitGroup, m *sync.Mutex, word string) {
	m.Lock()
	fmt.Println(word)
	m.Unlock()
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go say(&wg, &m, "alice: ping")
		wg.Wait()
		wg.Add(1)
		go say(&wg, &m, "bob: pong")
		wg.Wait()
	}
}
