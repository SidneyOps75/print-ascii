package main

import (
	"ascii-art/ascii-art"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 2 || lenArgs > 3 {
		fmt.Println("Incorrect no. of arguments. Expects: \"go run . <string> <banner name>\"")
		return
	}

	words := os.Args[1]
	str := strings.Split(words, "\\n")

	_, err := ascii.IsPrintableAscii(words)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := "standard.txt"
	if lenArgs == 3 {
		fileName = os.Args[2] + ".txt"
	}
	errCheck := ascii.ValidFile(fileName)
	if errCheck != nil {
		fmt.Println(errCheck)
		return
	}

	filePath := os.DirFS("./banner")
	content, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentString := string(content[1:])
	if fileName == "thinkertoy.txt" {
		contentString = strings.ReplaceAll(string(content[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")

	for _, word := range str {
		if word == "" {
			fmt.Println()
		} else {
			printAscii(word, contentSlice, 0)
		}

	}

}

func printAscii(word string, contentSlice []string, index int) {
	if index == 8 {
		return
	}

	for _, char := range word {

		character := contentSlice[int(char)-32]
		character = strings.ReplaceAll(character, "\r\n", "\n")
		lines := strings.Split(character, "\n")
		fmt.Printf(lines[index])
	}

	fmt.Println()
	printAscii(word, contentSlice, index+1)
}
