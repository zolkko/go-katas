package utils

import "strings"

func TrimInput(in string) string {
	in = strings.TrimSpace(in)
	in = strings.ToLower(in)
	return in
}
