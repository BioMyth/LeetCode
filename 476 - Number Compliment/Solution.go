package p476

import "math"

func findComplement(num int) int {
	mask := 0x1
	mask <<= uint(math.Log2(float64(num))) + 1
	mask -= 1
	return num ^ mask
}
