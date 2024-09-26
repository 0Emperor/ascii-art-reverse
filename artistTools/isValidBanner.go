package ascii

func IsValidBanner(args string) bool {
	if args == "standard" || args == "shadow" || args == "thinkertoy" || args == "standard.txt" || args == "shadow.txt" || args == "thinkertoy.txt" {
		return true
	}
	return false
}
