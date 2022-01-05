package main

import "math"

func TrappingRainwater(arr []int) int {

	var (
		maxL          = 0
		currentWater  = 0
		currentHeight = 0
	)

	p1 := 0
	p2 := len(arr) - 1
	maxR := arr[len(arr)-1]

	for p1 < p2 {
		if arr[p1] <= arr[p2] {
			if arr[p1] >= maxL {
				maxL = arr[p1]
			} else {
				currentHeight = arr[p1]
				currentWater += int(math.Min(float64(maxL), float64(maxR))) - currentHeight
			}
			p1++
		} else {
			if arr[p2] > maxR {
				maxR = arr[p2]
			} else {
				currentHeight = arr[p2]
				currentWater += int(math.Min(float64(maxL), float64(maxR))) - currentHeight
			}
			p2--
		}
	}

	return currentWater
}
