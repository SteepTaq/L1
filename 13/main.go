package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("a = %d, b = %d\n", a, b)
	swapXOR(&a, &b)
	fmt.Printf("после XOR обмена: a = %d, b = %d\n", a, b)

}

func swapXOR(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}
