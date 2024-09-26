package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/artistTools"
)

var ColorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	//"ColorSTOP": "\033[0m",
}

// main is the entry point of the program.
// It parses command-line arguments and generates ASCII art based on the provided inputs.
// The generated ASCII art is then outputted to the console or saved to a file.
func main() {
	args := len(os.Args[1:])
	Result := ""

	if args == 0 || args > 4 {
		printError()
		return
	}
	if strings.HasPrefix(os.Args[1], "--reverse=") {
	if args!=1 {
		printError()
		return
	}
		ascii.Reverse(os.Args[1][10:])
	return
	}
	template := "standard"
	color := ColorMap["white"]
	var fileName, substring string

	switch {
	case args == 1:
		if !ascii.IsValidASCII(os.Args[1]) {
			fmt.Println("Invalid ASCII characters in input")
			return
		}
		input := os.Args[1]
		Result = ascii.Artist(input, template, color, "")
		ascii.OutputFinal(Result, fileName)

	case args == 2:
		if ascii.IsOutputFlag(os.Args[1]) {
			if len(os.Args[1]) < 14 || !strings.HasSuffix(os.Args[1], ".txt") {
				printError()
				return
			}
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			input := os.Args[2]
			fileName = os.Args[1][9:]
			if !ascii.IsValidOutputFileName(fileName) {
				printError()
				return
			}
			Result = ascii.Artist(input, template, color, "")
			ascii.OutputFinal(Result, fileName)
		} else if ascii.IsColorFlag(os.Args[1]) {
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			input := os.Args[2]
			color := ColorMap[ascii.IsColor(os.Args[1][8:])]
			Result = ascii.Artist(input, template, color, "")
			Result = ascii.ApplyColor(Result, color)
			ascii.OutputFinal(Result, fileName)
		} else {
			if !ascii.IsValidASCII(os.Args[1]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[2]) {
				printError()
				return
			}
			template = Template(os.Args[2])
			input := os.Args[1]
			Result = ascii.Artist(input, template, color, "")
			ascii.OutputFinal(Result, fileName)
		}

	case args == 3:
		if ascii.IsOutputFlag(os.Args[1]) {
			if len(os.Args[1]) < 14 || !strings.HasSuffix(os.Args[1], ".txt") {
				printError()
				return
			}
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[3]) {
				printError()
				return
			}
			fileName = os.Args[1][9:]
			if !ascii.IsValidOutputFileName(fileName) {
				printError()
				return
			}
			input := os.Args[2]
			template = Template(os.Args[3])
			Result = ascii.Artist(input, template, color, substring)
			ascii.OutputFinal(Result, fileName)
		} else if ascii.IsColorFlag(os.Args[1]) {
			if ascii.IsValidSubString(os.Args[3], os.Args[2]) && !ascii.IsValidBanner(os.Args[3]) {
				if !ascii.IsValidASCII(os.Args[3]) {
					fmt.Println("Invalid ASCII characters in input")
					return
				}
				input := os.Args[3]
				substring = os.Args[2]
				color := ColorMap[ascii.IsColor(os.Args[1][8:])]
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			} else if ascii.IsValidBanner(os.Args[3]) {
				if !ascii.IsValidASCII(os.Args[2]) {
					fmt.Println("Invalid ASCII characters in input")
					return
				}
				input := os.Args[2]
				template = Template(os.Args[3])
				Result = ascii.Artist(input, template, color, substring)
				color := ColorMap[ascii.IsColor(os.Args[1][8:])]
				Result = ascii.ApplyColor(Result, color)
				ascii.OutputFinal(Result, fileName)
			} else {
				printError()
				return
			}
		} else {
			printError()
			return
		}

	case args == 4:
		if ascii.IsColorFlag(os.Args[1]) {
			color := ColorMap[ascii.IsColor(os.Args[1][8:])]
			if !ascii.IsValidASCII(os.Args[3]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[4]) {
				printError()
				return
			}
			if ascii.IsValidSubString(os.Args[3], os.Args[2]) {
				input := os.Args[3]
				template = Template(os.Args[4])
				substring = os.Args[2]
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			} else {
				input := os.Args[3]
				template = Template(os.Args[4])
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			}
		} else {
			printError()
			return
		}
	}
}

// Template generates a banner based on the provided template.
// If the template has a ".txt" extension, it removes the extension from the banner.
// Otherwise, it uses the template as is.
func Template(template string) string {
	banner := ""
	if strings.HasSuffix(template, ".txt") {
		banner = template[0 : len(os.Args[2])-4]
	} else {
		banner = template
	}
	return banner
}

func printError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}
