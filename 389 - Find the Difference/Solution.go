package p389

import "strings"

func findTheDifference(s string, t string) byte {
	for len(t) > 0 {
		prev := len(s)
		s = strings.Replace(s, string(t[0]), "", 1)
		if len(s) == prev {
			break
		}
		t = t[1:]
	}
	return t[0]
}
