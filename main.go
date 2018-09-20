package main

import (
	"fmt"
	"math/rand"

	"github.com/ferrandinand/go-testing-algoritms/algoritms"
)

func main() {
	a := RandomArray(150000)

	fmt.Print("Initial array:\n")
	fmt.Print(a)
	fmt.Print("\n")

	r := NoRepeated(a)
	fmt.Print(r)
}

// array with numbers  returns a arry without repeated numbers
func NoRepeated(numbersInput []int) []int {

	numbersInput = algoritms.Quick_sort(numbersInput)

	r := make([]int, 0, len(numbersInput))
	r = append(r, numbersInput[0])

	for _, number := range numbersInput {
		rTemp := r
		searchResult := algoritms.BinarySearch(rTemp, number)

		//Add number if not present
		if searchResult == -1 {
			r = append(r, number)
		}
	}
	return r
}

func RandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
