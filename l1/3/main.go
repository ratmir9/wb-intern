package main

import (
	"fmt"
	"sync"
)

// расчитывает и записывает квадрат числа в канал
func calculateSquare(num int, wg *sync.WaitGroup, sumChan chan<- int) {
	defer wg.Done()

	square := num * num
	sumChan <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sumChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg, sumChan)
	}

	wg.Wait()
	close(sumChan) // закрываем канал

	// находим сумму квадратов
	sum := 0
	for square := range sumChan {
		sum += square
	}

	fmt.Println("Сумма квадратов:", sum)
}
