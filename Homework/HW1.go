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
		fmt.Println("main: Hello")
		wg.Add(1)
		go say(&wg, &m, "goroutine: World")
		wg.Wait()
	}
}
