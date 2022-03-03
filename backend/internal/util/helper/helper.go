package helper

import "strings"

// ConstructPath takes a slice of tokens and construct
// a full path eg. 'slides/abc/dbe.txt', where the tokens
// are ["slides", "abc", "dbe.txt"].
func ConstructPath(tokens ...string) string {
	var sb strings.Builder
	for i, token := range tokens {
		if i == len(tokens)-1 {
			sb.WriteString(token)
		} else {
			sb.WriteString(token + "/")
		}
	}
	return sb.String()
}
