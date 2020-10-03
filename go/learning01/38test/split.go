package split

import "strings"

func S(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > - 1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}

	result = append(result, s)
	return
}

