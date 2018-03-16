package p520

import "unicode"

func detectCapitalUse(word string) bool {
	Valid := true
	HitLowers := false
	for index, val := range word {
		if unicode.IsLower(val) {
			if !HitLowers && index > 1 {
				Valid = false
			}
			HitLowers = true
		} else if index != 0 && HitLowers {
			Valid = false
		}
	}
	return Valid
}
