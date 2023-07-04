package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second) // Отправка значения каждую секунду
	}
}

func receiver(ch <-chan int, done chan<- bool) {
	for value := range ch {
		fmt.Println("Принято значение:", value)
	}

	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	N := 5 // Количество секунд работы программы

	go sender(ch)
	go receiver(ch, done)

	time.Sleep(time.Duration(N) * time.Second)
	close(ch)

	<-done
	fmt.Println("Программа завершена")
}
