package main

import (
	"fmt"
	"sync"
)

// рассчитывает квадрат числа
func calculateSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := num * num
	fmt.Println(square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10} // исходные данные
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg)
	}

	wg.Wait()
}
