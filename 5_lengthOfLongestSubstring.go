package main

import "math"

//t: O(n¨2)
//M: O(n)

func LengthOfLongestSubstring(str string) int {
	if str == "" {
		return 0
	}

	L := 0
	R := 0
	longest := 0
	currentLength := 0

	subs := make(map[string]int)
	for R < len(str) {
		val, ok := subs[string(str[R])]
		if ok && val >= L {
			currentLength = R - L
			longest = int(math.Max(float64(currentLength), float64(longest)))
			L = val + 1
		} else {
			subs[string(str[R])] = R
			R++
		}
	}
	currentLength = R - L
	longest = int(math.Max(float64(currentLength), float64(longest)))
	return longest
}
