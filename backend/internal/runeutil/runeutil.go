package runeutil

import "strings"

// Validator describes a function that validates a rune.
type Validator func(rune) bool

// AllowRunes creates a rune validator function from the given list of runes
// that will make it return true.
func AllowRunes(runes ...rune) Validator {
	set := make(map[rune]bool, len(runes))
	for _, r := range runes {
		set[r] = true
	}

	return func(r rune) bool {
		return set[r]
	}
}

// ContainsIllegal returns true if any of the runes in str matches none of the
// legal rune validators.
func ContainsIllegal(str string, legalRunes []Validator) bool {
	return strings.IndexFunc(str, func(r rune) bool {
		for _, fn := range legalRunes {
			if fn(r) {
				return false
			}
		}
		return true
	}) != -1
}
