package main

func isValid(s string) bool {
	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	if len(s)&1 == 1 {
		return false
	}

	myStack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		_, ok := pair[s[i]]
		if ok {
			myStack = append(myStack, pair[s[i]])
		} else if len(myStack) == 0 || s[i] != myStack[len(myStack)-1] {
			return false
		} else {
			myStack = myStack[:len(myStack)-1]
		}
	}

	return len(myStack) == 0
}
