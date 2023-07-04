package main

import "fmt"

func intersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	intersection := make([]int, 0)

	// Заполнение карты элементами первого множества
	for _, item := range set1 {
		setMap[item] = true
	}

	// Проверка элементов второго множества на присутствие в карте
	for _, item := range set2 {
		if setMap[item] {
			intersection = append(intersection, item)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	result := intersection(set1, set2)

	fmt.Println(result)
}
