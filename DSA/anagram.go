package languagebasics

// find the first indices in s where its the anagram of p
// Anagram : 2 strings are anagram where their frequencies are equal
// Approach : one way to use the count map of chars and compare it every time ( n* k), but we can improve
// by using the var matches to make comparison in o(1)

func AnagramIndices(s string, p string) []int {
	result := []int{}
	a := []rune(s)
	b := []rune(p)
	if len(a) < len(b) || len(b) == 0 {
		return result
	}

	window := make(map[rune]int)
	target := make(map[rune]int)
	for _, val := range b {
		target[val]++
	}

	for i, val := range a {
		window[val]++

		if i >= len(b) {
			leftChar := a[i-len(b)]
			window[leftChar]--
			if window[leftChar] == 0 {
				delete(window, leftChar)
			}
		}
		if i >= len(b)-1 && mapsEqual(window, target) {
			result = append(result, i-len(b)+1)
		}
	}

	return result
}

func mapsEqual[K comparable, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}

	return true
}
