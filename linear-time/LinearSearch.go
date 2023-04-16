package main

import "fmt"

func LinearSearch(haystack []int, needle int) bool {
	var iterCounter = 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			fmt.Printf("needle found. iterations taken: %d\n", i)
			return true
		}
		iterCounter++
	}

	fmt.Printf("needle not found. iterations taken: %d\n", iterCounter)
	return false
}

func main() {
	var haystack = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	LinearSearch(haystack, 6)
	LinearSearch(haystack, 15)

}
