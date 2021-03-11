package main

func calculate(s string) int {
	myStack := make([]int, 0)
	tempNumber := 0
	var preFlag byte = '+'
	i := 0
	if s[0] == '-' {
		preFlag = '-'
		i = 1
	}

	for ; i < len(s); i++ {
		isDigitFlag := '0' <= s[i] && s[i] <= '9'
		if isDigitFlag {
			tempNumber = tempNumber*10 + int(s[i]-'0')
		}
		if !isDigitFlag && s[i] != ' ' || i == len(s)-1 {
			switch preFlag {
			case '+':
				myStack = append(myStack, tempNumber)
			case '-':
				myStack = append(myStack, -tempNumber)
			case '*':
				myStack[len(myStack)-1] *= tempNumber
			case '/':
				myStack[len(myStack)-1] /= tempNumber
			}
			preFlag = s[i]
			tempNumber = 0
		}
	}

	sumResult := 0
	for _, each := range myStack {
		sumResult += each
	}

	return sumResult
}

// 每次遇到一个符号（+-*/）时，才能够判断上一个符号的运算方式，例如11+2*3+5+6-7
// 遇到*时，判断上一个+后面的2可以直接入栈，保留preFlag=*；而当遇到第2个+时，判断上一个*需要计算结果了
// 所以计算myStack[len(myStack)-1] *= 3，保留preFlag=+，当遇到第3个+时，判断上一个+，入栈5，保留preFlag=+
// 当遇到最后一个-时，判断上一个+，入栈6，保留preFlag=-；
// 7所对应的i为len(myStack)-1，所以直接取负数入栈，此使preFlag被设置成什么（最后一个字符是数字）已经不重要，因为已经跳出循环了
