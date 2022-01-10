package string

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		if strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
