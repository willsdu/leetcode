package primary

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
*/
func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		exchange := l - 1 - i
		s[i], s[exchange] = s[exchange], s[i]
	}
}

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func reverse(x int) int {
	min := (1 << 31) * -1
	max := (1 << 31) - 1

	resullt := 0
	for {
		if x == 0 {
			break
		}
		mod := x % 10
		resullt = resullt*10 + mod
		x = x / 10
		if resullt > max || resullt < min {
			return 0
		}
	}
	return resullt
}

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
*/
func firstUniqChar(s string) int {
	smap := make(map[rune]int)
	for k, v := range s {
		_, ok := smap[v]
		if ok {
			smap[v] = -1
		} else {
			smap[v] = k
		}
	}
	min := -1
	for _, v := range smap {
		if min == -1 && v != -1 {
			min = v
		}
		if v != -1 && v < min {
			min = v
		}
	}
	return min
}
