package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := make(map[string]int)
	mutex := sync.Mutex{}

	wg.Add(3)

	go func() {
		defer wg.Done()

		mutex.Lock()
		data["apple"] = 10
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()

		mutex.Lock()
		data["banana"] = 20
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()

		mutex.Lock()
		data["cherry"] = 30
		mutex.Unlock()
	}()

	wg.Wait()

	fmt.Println(data)
}
