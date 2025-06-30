package main

import "fmt"


func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true
		}
		m[num] = true
	}
	return false
}