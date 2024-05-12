package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"ascii-art/ascii-art"
)

func main() {
	lenArgs := len(os.Args)
	//Check if the number of arguments is incorrect
	if lenArgs < 2 || lenArgs > 3 {
		fmt.Printf("Incorrect no. of arguments.\nExpects: \"go run . <string> | cat -e\"\nor\n\"go run . <string> <banner name> | cat -e\"\n")
		return
	}
    // String argument from the command line
	word := os.Args[1]
	str := strings.Split(word, "\\n")
    // Call a function to check if the characters in the string are printable ascii characters
	_, err := ascii.IsPrintableAscii(word)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := "standard.txt" //Default filename
	if lenArgs == 3 {
		// Override default filename if the filename is provided in the command line
		fileName = os.Args[2] + ".txt"
	}
	// the ValidFile function checks if the filename is valid
	errCheck := ascii.ValidFile(fileName)
	if errCheck != nil {
		fmt.Println(errCheck)
		return
	}
    // The file system path and reading of the content of the filename
	filePath := os.DirFS("./banner")
	content, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Converting the content bytes into a string
	contentString := string(content[1:])
	// Special handling for thinkertoy
	if fileName == "thinkertoy.txt" {
		contentString = strings.ReplaceAll(string(content[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n") // split the contentstring into slices using double new lines
    
	for _, word := range str{
		if word == ""{
			fmt.Println()
		} else{
	//Call the printAscii function to print the ascii characters
	PrintAscii(word, contentSlice, 0)
		}
}
}
// This function is used to print the ascii characters
func PrintAscii(word string, contentSlice []string, index int) {
	if index == 8 { //The base case returns nothing when it reaches index 8
		return
	}
    
	for _, char := range word { // Looping through each character in the input string.
		character := contentSlice[int(char)-32] // Get the ASCII art representation of the character.
		character = strings.ReplaceAll(character, "\r\n", "\n") //Replace the carriage return with newline
		lines := strings.Split(character, "\n") //Split ascii art into lines
		fmt.Printf(lines[index]) //Print the line corresponding to the current index
	}
	fmt.Println()// Print newline after printing all characters
	PrintAscii(word, contentSlice, index+1)// recursively call the function with a increment of the index.
}
