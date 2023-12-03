package main

func lengthOfLongestSubstring(s string) int {
	longest := 0
	start := 0
	m := make(map[rune]int)
	for index, character := range s {
		if val, ok := m[character]; ok {
			if start < val+1 {
				start = val + 1
			}
		}
		m[character] = index
		longest = max(longest, index-start+1)
	}
	return longest
}
