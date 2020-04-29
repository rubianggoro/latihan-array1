package main

import "fmt"

func difference(arr1 [20]int, arr2 [20]int) []int {

	var diff []int
	for i := 0; i < 2; i++ {

		for _, value1 := range arr1 {
			f := false

			for _, value2 := range arr2 {
				if value1 == value2 {
					f = true
				}
			}
			if !f {
				diff = append(diff, value1)
			}
		}
		if i == 0 {
			arr1, arr2 = arr2, arr1
		}
	}
	return diff
}

func main() {
	a := [20]int{4, 5, 13, 35, 76, 83, 16, 59, 10, 78, 87, 11, 25, 90, 38, 11, 19, 17, 82, 12}
	b := [20]int{4, 5, 6, 76, 8, 45, 90, 78, 10, 98, 18, 35, 13, 5, 35, 38, 11, 83, 4, 16}

	fmt.Println(difference(a, b))
}
