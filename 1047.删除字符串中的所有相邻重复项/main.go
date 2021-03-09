package main

func removeDuplicates(S string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(S); i++ {
		if len(stack) == 0 || S[i] != stack[len(stack)-1] {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
