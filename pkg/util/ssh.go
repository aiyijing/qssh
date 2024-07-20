package util

import "strings"

func ParseSSHURL(sshURL string) (user string, host string) {
	pairs := strings.SplitN(sshURL, "@", 2)
	if len(pairs) != 2 {
		return "", pairs[0]
	}
	return pairs[0], pairs[1]
}
