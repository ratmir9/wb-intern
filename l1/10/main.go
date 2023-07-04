package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		key := int(math.Round(temp/10.0)) * 10
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupedTemperatures := groupTemperatures(temperatures)

	for key, value := range groupedTemperatures {
		fmt.Printf("%d: %v\n", key, value)
	}
}
