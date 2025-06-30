package main

import "fmt"

func main() {

	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result) // Output: 3, because "abc" is the longest substring without repeating characters
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		if index, found := m[s[right]]; found && index >= left {
			left = index + 1
		}
		m[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}