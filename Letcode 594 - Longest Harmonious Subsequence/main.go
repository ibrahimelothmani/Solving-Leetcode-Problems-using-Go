package main

import "fmt"


func main() {
	nums := []int{1, 3, 2, 2, 5, 4, 3, 2}
	result := findLHS(nums)
	fmt.Println("Length of the longest harmonious subsequence:", result) // Output: 5
}

func findLHS(nums []int) int {
	// Create a map to store the frequency of each number.
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	maxLength := 0
	// Iterate through the map to find harmonious pairs.
	for num, count := range freqMap {
		// Check if a number with a difference of 1 exists.
		if nextCount, ok := freqMap[num+1]; ok {
			currentLength := count + nextCount
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}
