package ascii

import (
	"os"
)

// ReadFile reads the contents of a file and returns them as a slice of strings.
// Each string in the slice represents a line in the file.
// The file is read line by line, where each line is terminated by a newline character ('\n').
// If an error occurs while reading the file, it is returned along with a nil slice.
// ELse it return content and nil error.
func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var content []string
	singleByte := make([]byte, 1)
	line := ""
	for {
		_, err = file.Read(singleByte)
		if err != nil {
			break
		}
		if singleByte[0] == '\n' {
			content = append(content, line)
			line = ""
		} else {
			line += string(singleByte)
		}
	}
	content = append(content, line)
	// content = append(content, "")
	return content, nil
}
