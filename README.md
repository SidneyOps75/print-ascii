# ASCII Art Generator

## Overview

The ASCII Art Generator is a command-line tool written in Go that converts text input into ASCII art representations. It reads ASCII art templates from text files and prints the corresponding ASCII art for each character in the input string.

## Features

- Converts text input into ASCII art representations.
- Supports a variety of ASCII art templates stored in text files.
- Handles input validation to ensure only valid characters are used.
- Flexible command-line interface allows customization of input and template files.

## Usage

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/print-ascii.git

2. Navigate to the directory
   ```bash
   cd ascii-art

### Running the program

1.  To generate the ASCII art run the following command in the first and second argument
    ```bash
    go run . <word string> <filename> 


### Filename

The ASCII art filenames are stored as text files in the banner directory. Each template file contains ASCII art representations for characters in the ASCII range 32 to 127.


### Error Handling

1. If the specified template file does not exist, the program will print an error message.

2. If any character in the input word falls outside the range of ASCII characters (32 to 127), the program will print an error message.



