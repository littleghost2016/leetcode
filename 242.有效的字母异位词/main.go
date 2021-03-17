package main

func isAnagram(s string, t string) bool {
	hashArray := [26]int{}
	if len(s) != len(t) {
		return false
	}
	length := len(s)

	for i := 0; i < length; i++ {
		hashArray[int(s[i]-'a')]++
	}

	for i := 0; i < length; i++ {
		hashArray[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if hashArray[i] != 0 {
			return false
		}
	}

	return true
}

// 写了一次没改就成功，嘿嘿
