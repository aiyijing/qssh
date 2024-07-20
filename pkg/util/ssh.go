package util

import "strings"

func ParseConnArgs(connArgs string) (user string, host string) {
	pairs := strings.SplitN(connArgs, "@", 2)
	if len(pairs) != 2 {
		return "", pairs[0]
	}
	return pairs[0], pairs[1]
}
