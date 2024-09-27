package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	}
	if len(phrases) == 2 {
		return strings.Join(phrases, " and ")
	} else if len(phrases) == 1 {
		return phrases[0]
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]

	return result
}
