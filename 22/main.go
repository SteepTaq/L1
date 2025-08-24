package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 26)
	b := big.NewInt(1 << 30)
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Printf("a+b = %s\n", sum.String())
	fmt.Printf("a-b = %s\n", diff.String())
	fmt.Printf("a*b =%s\n", mul.String())
	fmt.Printf("a/b = %s\n", div.String())
}
