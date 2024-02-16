package main

import (
	"fmt"
)

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	dst := IncrementOdd
	newFunction := appendFunc(dst, RevereSlice, PrintSlice)
	newFunction(slice)
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, fn := range src {
			fn(slice)
		}
	}
}

func IncrementOdd(slice []int) {
	for i := 0; i < len(slice); i++ {
		if i%2 == 1 {
			slice[i]++
		}
	}
}

func PrintSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
}

func RevereSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < len(slice)/2 && j > len(slice)/2; i, j = i+1, j-1 {
		temp := slice[i]
		slice[i] = slice[j]
		slice[j] = temp
	}
}
