package main

import (
	"math"
)

func Max_area(arr []int) int {
	a := 0
	b := len(arr) - 1

	height := 0
	width := 0
	area := 0
	max_area := 0

	for i := 0; i < len(arr); i++ {
		height = int(math.Min(float64(arr[a]), float64(arr[b])))
		width = b - a
		area = height * width
		max_area = int(math.Max(float64(max_area), float64(area)))

		if arr[a] > arr[b] {
			b--
		} else {
			a++
		}
	}

	return max_area
}
