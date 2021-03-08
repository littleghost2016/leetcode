package main

func isSubsequence(s string, t string) bool {
	sLength, tLength := len(s), len(t)

	if sLength == 0 {
		return true
	}

	s_i := 0
	for t_i := 0; t_i < tLength; t_i++ {
		if s[s_i] == t[t_i] {
			s_i++
			if s_i == sLength {
				return true
			}
		}
	}

	return false
}
