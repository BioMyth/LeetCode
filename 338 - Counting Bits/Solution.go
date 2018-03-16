package p338

import "math"

func countBits(num int) []int {
	var ret []int
	ret = append(ret, 0)
	for i := 1; i <= num; i++ {
		tmp := 0
		for bit := uint(0); float64(bit) < math.Log2(float64(i))+1; bit++ {
			tmp += (i >> bit) & 0x1
		}
		ret = append(ret, tmp)
	}
	return ret
}
