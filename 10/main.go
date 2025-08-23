package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemp(temp)

	fmt.Println("группировка температурных колебаний:")
	for groupKey, values := range groups {
		fmt.Printf("%.0f: %v\n", groupKey, values)
	}
}

func groupTemp(temps []float64) map[float64][]float64 {
	groups := make(map[float64][]float64)
	for _, temp := range temps {
		groupKey := math.Floor(temp/10) * 10
		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}
