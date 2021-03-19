package main

func canConstruct(ransomNote string, magazine string) bool {
	hashArray := [26]int{}
	for i := 0; i < len(magazine); i++ {
		hashArray[int(magazine[i]-'a')]++
	}
	for i := 0; i < len(ransomNote); i++ {
		hashArray[int(ransomNote[i]-'a')]--
		if hashArray[int(ransomNote[i]-'a')] < 0 {
			return false
		}
	}

	return true
}
