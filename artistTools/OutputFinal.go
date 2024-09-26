package ascii

import (
	"fmt"
	"os"
)

// OutputFinal print our FInal Ascii Graph either in console or in file
func OutputFinal(result string, fileName string) {
	if fileName != "" {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error Creating file")
			return
		}
		os.WriteFile(fileName, []byte(result), 0644)
	} else {
		fmt.Print(result)
	}
}
