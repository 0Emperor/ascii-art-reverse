package ascii

// check if our input is valid ASCII and doesn't contain invalid chars
// IsValidASCII checks if a given word contains only valid ASCII characters.
// It returns true if all characters in the word are within the ASCII range of 32 to 126 (inclusive),
// and false otherwise.
func IsValidASCII(word string) bool {
	for _, r := range word {
		if !(r >= 32 && r <= 126) {
			return false
		}
	}
	return true
}
