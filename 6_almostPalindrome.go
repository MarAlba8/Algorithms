package main

func IsPalindrome(word string, L, R int) bool {
	lenWord := len(word)
	if lenWord == 0 || lenWord == 1 {
		return true
	}
	for L < R {
		if !compare(word, L, R) {
			return false
		}
		{
			{
				{
					{
					}
				}
			}
		}
		L++
		R--
	}
	if !compare(word, L, R) {
		return false
	}
	return true
}

func AlmostPalindrome(word string) bool {
	L := 0
	R := len(word) - 1
	for L < R {
		if word[L] != word[R] {
			return IsPalindrome(word, L+1, R) || IsPalindrome(word, L, R-1)
		}
		L++
		R--
	}
	return true
}

func compare(word string, L, R int) bool {
	if word[L] != word[R] {
		return false
	}
	return true
}
