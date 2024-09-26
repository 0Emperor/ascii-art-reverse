package ascii

import (
	"fmt"
	"os"
	"strings"
)

var ColorStop = "\033[0m"

// Artist generates ASCII art based on the provided input, template, color, and substring.
// It reads the ASCII art template from a file, replaces any carriage return characters,
// splits the input into lines, and generates ASCII art for each line using the LineAsAscii function.
// The generated ASCII art is then concatenated and returned as a single string.
func Artist(input, template, color, substring string) string {
	// check for errors too in case file no longer exist
	asciiGraph, err := ReadFile("banners/" + template + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	for idx, LINE := range asciiGraph {
		asciiGraph[idx] = strings.ReplaceAll(LINE, "\r", "")
	}
	lines := strings.Split(input, "\\n")
	// If the input string is empty after removing "\\n", it removes the first line.
	if strings.ReplaceAll(input, "\\n", "") == "" {
		lines = lines[1:]
	}
	finalAsciiArt := ""
	for _, line := range lines {
		// add each line to finalAsciiArt using the LineAsAscii function.
		finalAsciiArt += LineAsAscii(line, color, substring, asciiGraph)
	}
	return finalAsciiArt
}

// LineAsAscii generates an ASCII art representation of a given line with color highlighting for a specified substring.
// It takes the line to be converted, the color for highlighting, the substring to be highlighted, and the ASCII graph as input.
// The function returns the final ASCII art representation as a string.
func LineAsAscii(line, color, substring string, asciiGraph []string) string {
	var asciiChars []string
	finalAsciiArt := ""
	// substring := "X"
	substringIndex := 0
	if line != "" {
		if strings.Contains(line, substring) {
			substringIndex = strings.Index(line, substring)
		}
		for idx, char := range line {
			if idx >= substringIndex && idx < substringIndex+len(substring) {
				for i := 8; i > 0; i-- {
					asciiChars = append(asciiChars, color+string(asciiGraph[findLastLine(char)-i])+ColorStop)
				}
			} else {
				for i := 8; i > 0; i-- {
					asciiChars = append(asciiChars, string(asciiGraph[findLastLine(char)-i]))
				}
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiChars); j += 8 {
				finalAsciiArt += asciiChars[i+j]
			}
			finalAsciiArt += "\n"
		}
	} else {
		finalAsciiArt += "\n"
	}
	return finalAsciiArt
}

// ApplyColor applies the specified color to the given ASCII art.
// It returns the modified ASCII art with the color applied.
func ApplyColor(asciiArt, color string) string {
	return fmt.Sprintf("%s%s%s", color, asciiArt, ColorStop)
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
