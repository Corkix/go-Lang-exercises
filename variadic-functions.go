package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func varidicfunc() {
	sum(1, 2)
	sum(1, 2, 3)

	num := []int{1, 2, 3, 4}
	sum(num...)
}
