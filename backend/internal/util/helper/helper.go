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

// RemoveDuplicatesFromStrings takes a slice of string and remove
// duplicates in the array
func RemoveDuplicatesFromStrings(arr []string) []string {
	// initialize a hashmap
	hashmap := make(map[string]bool)
	// a temporary slice to store unduplicated element
	temp := []string{}

	// loop through the arr
	for _, elem := range arr {
		// if elem not in hashmap
		if _, ok := hashmap[elem]; !ok {
			// insert it and the element to the slice
			hashmap[elem] = true
			temp = append(temp, elem)
		}
	}

	return temp
}
