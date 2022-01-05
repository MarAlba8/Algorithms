package main

func typedOutString(str1, str2 string) bool {

	backCount := 0
	p1 := len(str1) - 1
	p2 := len(str2) - 1

	for p1 >= 0 || p2 >= 0 {
		if str1[p1] == '#' || str2[p2] == '#' {
			if str1[p1] == '#' {
				backCount = 2
				for backCount > 0 {
					p1--
					backCount--
					if p1 > 0 && str1[p1] == '#' {
						backCount += 2
					}
				}
			}

			if str2[p2] == '#' {
				backCount = 2
				for backCount > 0 {
					p2--
					backCount--
					if p2 > 0 && str2[p2] == '#' {
						backCount += 1
					}
				}
			}

			if str1[p1] != str2[p2] {
				return false
			} else {
				p1--
				p2--
			}
		}
	}
	return true
}
