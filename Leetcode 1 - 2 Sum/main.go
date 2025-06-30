package main

import "fmt"

func main() {
	
	nums := []int{5, 7, 11, 2}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}