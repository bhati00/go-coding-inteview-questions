package languagebasics

import "strings"

func ReverseString(s string) string {

	words := strings.Fields(s) // splits by any number of spaces
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
