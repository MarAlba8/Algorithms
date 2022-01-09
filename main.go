package main

import "fmt"

func main() {

	//Exercice 1
	//arr := []int{1, 3, 5, 8}
	//fmt.Println(find_addens(arr))

	//Exercices 2
	//arr := []int{4, 8, 1, 2, 3, 9}
	//fmt.Println(Max_area(arr))

	//Exercices 3
	//arr := []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}
	//fmt.Println(TrappingRainwater(arr))

	//Exercices 4
	//str1 := "ab##"
	//str2 := "c#d#"
	//fmt.Print(typedOutString(str1, str2))

	//Exercices 5
	//str := "abcbda"
	//fmt.Print(LengthOfLongestSubstring(str))

	//Exercices 6
	//str = "A man, a plan, a canal: Panama"
	//str := "cameleon"
	//fmt.Println(AlmostPalindrome(str))

	//Ex linked list

	//head := CreateLinkedList(5)
	//fmt.Println("After reverse")

	//Question 7
	//newList := ReversePortion(head, 1, 2)
	//PrintListValues(newList)

	//Question 8
	//Create DlinkedList
	head := create_dlist_with_child(6, 1)
	PrintDListValues(head)
	//Solution
	head = FlatList(head)
	fmt.Println("After flatting")
	PrintDListValues(head)
}

var parents_index = map[int]int{
	3: 4,
	8: 2,
}
