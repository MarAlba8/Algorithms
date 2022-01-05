package main

func find_addens(arr []int, target int) []int {
	var number_to_find int
	numbers_to_find := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		number_to_find = target - arr[i]
		found, ok := numbers_to_find[arr[i]]
		if ok {
			addens := []int{found, i}
			return addens
		}
		numbers_to_find[number_to_find] = i
	}
	return nil
}
