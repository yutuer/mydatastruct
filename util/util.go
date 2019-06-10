package util

func ToPow2(i int) int {
	if i < 1<<3 {
		return 1 << 3
	} else if i&(i-1) == 0 {
		return i
	} else {
		var j uint
		// 求最左边的1是第几位
		for ; i > 0; i >>= 1 {
			j++
		}
		return 1 << j
	}
}
