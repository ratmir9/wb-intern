package main

import (
	"fmt"
)

func setBit(n int64, pos uint, bitValue int) int64 {
	mask := int64(1 << pos) // Создание маски с установленным битом в позиции pos

	if bitValue == 1 {
		n |= mask // Установка бита в 1
	} else {
		n &= ^mask // Установка бита в 0
	}

	return n
}

func main() {
	var num int64 = 24 // Пример значения int64
	pos := uint(3)     // Позиция бита для установки (начиная с 0)
	bitValue := 1      // Значение бита (0 или 1)

	result := setBit(num, pos, bitValue)
	fmt.Printf("Результат: %d\n", result)
}
