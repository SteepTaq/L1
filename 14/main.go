package main

import "fmt"

func main() {
	whatType(52)
	whatType("wb")
	whatType(false)
	whatType(make(chan int))
	whatType(3.14)
}

func whatType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Тип: int, значение: %v\n", v)
	case string:
		fmt.Printf("Тип: string, значение: %v\n", v)
	case bool:
		fmt.Printf("Тип: bool, значение: %v\n", v)
	case chan int:
		fmt.Printf("Тип: chan int, значение: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T, значение: %v\n", v, v)
	}
}
