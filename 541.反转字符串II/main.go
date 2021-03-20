package main

func reverseStr(s string, k int) string {

	// go不能直接对字符串进行按照索引位置更改内容（因为字符串是不可变的）
	// 强行更改会报错 cannot assign to s[0] (type of value byte)
	// 因此要先转换成[]byte
	ss := []byte(s)

	// 这里i+=(2*K)控制的是步进
	for i := 0; i < len(ss); i += (2 * k) {
		if i+k <= len(ss) {
			// 注意下面的reverse函数实现的是左闭右闭，所以i+k要再-1
			reverse(ss, i, i+k-1)
		} else {
			// 注意这里一直到最后，所以 不需要 在len(ss)-1的基础上再-1
			reverse(ss, i, len(ss)-1)
		}
	}

	return string(ss)
}

// 左闭右闭
func reverse(ss []byte, left, right int) {
	for left < right {
		ss[left], ss[right] = ss[right], ss[left]
		left++
		right--
	}
}
