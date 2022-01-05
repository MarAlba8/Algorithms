package main

import (
	"log"
	"regexp"
	"strings"
)

func IsPalindrome1(word string) (bool, int, int) {
	lenWord := len(word)
	if lenWord == 0 || lenWord == 1 {
		return true, 0, 0
	}
	c := 0
	if lenWord%2 == 0 {
		c = 1
	}

	L := 0
	R := lenWord - 1
	for L+c < R {
		if !compare(word, L, R) {
			return false, L, R
		}
		L++
		R--
	}
	if !compare(word, L, R) {
		return false, L, R
	}
	return true, 0, 0
}

func IsPalindrome2(word string) bool {
	word = clearSentence(word)
	lenWord := len(word)
	if lenWord == 0 || lenWord == 1 {
		return true
	}
	L := 0
	R := 0
	if lenWord%2 == 0 {
		R = lenWord / 2
		L = R - 1
	} else {
		L = lenWord / 2
		R = lenWord / 2
	}

	for L >= 0 {
		if !compare(word, L, R) {
			return false
		}
		L--
		R++
	}
	return true
}

func IsPalindrome3(word string) (bool, int, int) {
	word = clearSentence(word)
	lenWord := len(word)
	if lenWord == 0 || lenWord == 1 {
		return true, -1, -1
	}
	L := 0
	R := 0
	if lenWord%2 == 0 {
		R = lenWord / 2
		L = R - 1
	} else {
		L = lenWord / 2
		R = lenWord / 2
	}

	for L >= 0 {
		if !compare(word, L, R) {
			return false, L, R
		}
		L--
		R++
	}
	return true, -1, -1
}

func clearSentence(sentence string) string {
	sentence = strings.ToLower(sentence)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newSentence := reg.ReplaceAllString(sentence, "")
	return newSentence
}
