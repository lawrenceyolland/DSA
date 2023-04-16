package main

import (
	"fmt"
	"math"
)

func BinarySearch(haystack []int, needle int) bool {
	var start int = 0
	var end int = len(haystack)

	for start < end {
		midPoint := int(math.Floor(float64(start) + (float64(end)-float64(start))/2))
		midValue := haystack[midPoint]

		if midValue == needle {
			return true
		} else if midValue > needle {
			end = midPoint
		} else {
			start = midPoint + 1
		}
	}

	return false
}

func main() {
	var haystack = []int{1, 2, 3, 4, 6, 7, 9, 10}
	var needle = 6

	res := BinarySearch(haystack, needle)
	fmt.Printf("is needle (%d) in haystack? %t", needle, res)
}
