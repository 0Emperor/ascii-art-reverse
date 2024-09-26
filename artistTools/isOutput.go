package ascii

import "strings"

func IsOutputFlag(flag string) bool {
	return strings.HasPrefix(flag, "--output=")
}

func IsValidOutputFileName(str string) bool {
	return len(str) > 4 && strings.HasSuffix(str, ".txt")
}
