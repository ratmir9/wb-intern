package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Before sleep")
	Sleep(2 * time.Second)
	fmt.Println("After sleep")
}
