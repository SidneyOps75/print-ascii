package ascii

import (
	"fmt"
	"os"
	"strings"
)

func IsPrintableAscii(str string) (bool, error) {
	var nonPritables string
	for _, char := range str {
		if char < ' ' || char > '~' {
			nonPritables += string(char)
		}
	}
	if nonPritables != "" {
		return false,
			fmt.Errorf("%s are not within the printable ascii range", nonPritables)
	}
	return true, nil
}

func ValidFile(fileName string) error {
	path := "./banner"
	openPath, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error")
	}
	defer openPath.Close()
	filenames, err := openPath.Readdirnames(0)
	if err != nil {
		return fmt.Errorf("error name")
	}
	fileNameString := strings.Join(filenames, " ")
	if !strings.Contains(fileNameString, fileName) {
		return fmt.Errorf("not a valid banner file name")
	}
	return nil
}
