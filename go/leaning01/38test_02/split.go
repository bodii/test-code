package split

import "strings"

func S(str, sep string) (result []string) {
	i := strings.Index(str, sep)

	for i > -1 {
		result = append(result, str[:i])
		str = str[i+1:]
		i = strings.Index(str, sep)
	}

	result = append(result, str)
	return
}
