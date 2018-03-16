package p136

func singleNumber(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret = ret ^ val
	}
	return ret
}
