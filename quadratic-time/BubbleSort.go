package main

import (
	"fmt"
)

func Swap(mySlice []int, idx int) {
	mySlice[idx], mySlice[idx+1] = mySlice[idx+1], mySlice[idx]
}

func BubbleSort(mySlice []int) {
	var iterCounter = 0
	var swapCounter = 0

	for i := 0; i < len(mySlice); i++ {
		for j := 0; j < len(mySlice)-1-i; j++ {
			if mySlice[j] > mySlice[j+1] {
				Swap(mySlice, j)
				swapCounter++
			}
			iterCounter++
		}
	}
	fmt.Printf("Sort Completed. Iterations: %d. Swaps made: %d. Final Slice: %d\n", iterCounter, swapCounter, mySlice)
}

func main() {
	var testOne = []int{6, 5, 10, 4, 2, 1, 8, 7, 3}
	BubbleSort(testOne)

	var testTwo = []int{1, 2, 3, -1, 5, -2, 7, 8, 10}
	BubbleSort(testTwo)

	var testThree = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}
	BubbleSort(testThree)
}
