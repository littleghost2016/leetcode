package main

func isOneBitCharacter(bits []int) bool {

	// 最后一个字符不是0而是1则直接false
	if bits[len(bits)-1] != 0 {
		return false
	}

	// 在上面的判断之后，长度为1和2的直接为true
	if len(bits) < 2 {
		return true
	}

	// 从第一个字符开始，遇到1则索引值+2，遇到0则索引值加1，一直到倒数第二个字符停止
	index := 0
	for index < (len(bits) - 1) {
		if bits[index] == 0 {
			index += 1
		} else {
			index += 2
		}

		// 以上五行代码可以用以下一行代替
		// 因为遇到0则+1，遇到1则+2
		// index += (bits[index]+1)
	}

	// 最后一个0是单独的，即和前面可能存在的1不会构成10或者11
	if index == (len(bits) - 1) {
		return true
	} else {
		return false
	}
}
