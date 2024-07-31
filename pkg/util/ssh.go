package util

import "strings"

func ParseSSHURL(sshURL string) (user string, host string) {
	pairs := strings.SplitN(sshURL, "@", 2)
	if len(pairs) != 2 {
		return "", pairs[0]
	}
	return pairs[0], pairs[1]
}

func ParseSSHURLWithPath(sshURL string) (user string, host string, remotePath string) {
	pairs := strings.SplitN(sshURL, ":", 2)
	if len(pairs) != 2 {
		return "", "", ""
	}
	remotePath = pairs[1]
	pairs = strings.SplitN(pairs[0], "@", 2)
	if len(pairs) != 2 {
		user = ""
		host = pairs[0]
		return
	}
	user = pairs[0]
	host = pairs[1]
	return
}
