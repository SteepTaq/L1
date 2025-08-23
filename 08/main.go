package main

import "fmt"

func main() {
	var num int64
	var position int
	var value int

	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &num)

	fmt.Print("Введите позицию бита: ")
	fmt.Scanf("%d", &position)

	fmt.Print("Введите значение бита(0 или 1): ")
	fmt.Scanf("%d", &value)

	fmt.Printf("Исходное число: %d\n", num)
	fmt.Printf("Двоичное: %016b\n", num)

	result := setBit(num, position, value)
	fmt.Printf("Результат: %d\n", result)
	fmt.Printf("Двоичное: %016b\n", result)

}

func setBit(num int64, position int, value int) int64 {
	if value == 1 {
		return num | (1 << position)
	} else {
		return num &^ (1 << position)
	}
}
