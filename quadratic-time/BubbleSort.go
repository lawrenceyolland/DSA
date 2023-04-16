package main

import (
	"fmt"
)

var counter int = 0

func resetCounter() {
	counter = 0
}

func Swap(mySlice []int, idx int) {
	mySlice[idx], mySlice[idx+1] = mySlice[idx+1], mySlice[idx]
}

func BubbleSort(mySlice []int) {
	var swapCounter int = 0

	for i := 0; i < len(mySlice)-1; i++ {
		if mySlice[i] > mySlice[i+1] {
			// mySlice[i], mySlice[i+1] = mySlice[i+1], mySlice[i]
			Swap(mySlice, i)
			swapCounter++
		}
		counter++
	}

	if swapCounter == 0 {
		fmt.Printf("Sort Completed. Iterations taken: %d. Final Slice: %d\n", counter, mySlice)
		return
	}

	BubbleSort(mySlice)
}

func main() {
	var testOne = []int{6, 5, 10, 4, 2, 1, 8, 7, 3}
	BubbleSort(testOne)
	resetCounter()

	var testTwo = []int{1, 2, 3, -1, 5, -2, 7, 8, 10}
	BubbleSort(testTwo)
	resetCounter()

	var testThree = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -1}
	BubbleSort(testThree)
	resetCounter()
}
