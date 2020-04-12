package main

import (
	"fmt"
	"sort"
)

func median(i []int) float64 {
	sort.Ints(i)
	var res float64
	if len(i)%2 == 0 {
		res = float64(i[len(i)/2]+i[len(i)/2-1]) / 2
	} else {
		res = float64(i[(len(i) / 2)])
	}
	return res
}

func main() {
	//I wants to make two examples
	//For first i use array consisting of 5, 3, 6, 8, 4, 9. In this case result must be 5.5
	test := []int{5, 3, 6, 8, 4, 9}
	fmt.Println("First example:", median(test))

	//Second example use array consisting of 4, 3, 7, 9, 15. The expected result is 7.
	test = []int{4, 3, 7, 9, 15}
	fmt.Println("Second example:", median(test))
}


