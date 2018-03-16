package p406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0]) || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	var ret [][]int
	for index := 0; index < len(people); index++ {
		ret = append(ret[:people[index][1]], append([][]int{people[index]}, ret[people[index][1]:]...)...)
	}
	return ret
}
