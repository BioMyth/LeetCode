package p500

import "strings"

func findWords(words []string) []string {
	ret := make([]string, 0, len(words))

	for _, s := range words {
		numIn := 0
		if strings.ContainsAny(s, "qwertyuiop") {
			numIn++
		}
		if strings.ContainsAny(s, "asdfghjkl") {
			numIn++
		}
		if strings.ContainsAny(s, "zxcvbnm") {
			numIn++
		}
		if numIn == 1 {
			ret = append(ret, s)
		}
	}
	return ret
}
