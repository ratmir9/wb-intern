package main

import "fmt"

func RemoveElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 10
	newSlice := RemoveElement(slice, index)
	fmt.Println(newSlice)
}
