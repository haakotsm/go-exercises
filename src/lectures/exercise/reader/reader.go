//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:

//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//* When the user enters either "hello" or "bye", the program
	//  should respond with a message after pressing the enter key.
	//* Whenever the user types a "Q" or "q", the program should exit.
	//* Upon program exit, some usage statistics should be printed
	//  ('Q' and 'q' do not count towards these statistics):
	//  - The number of non-blank lines entered
	//  - The number of commands entered
	r := bufio.NewReader(os.Stdin)

	commands := 0
	numberOfLines := 0

	for {
		input, inputErr := r.ReadString('\n')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		if strings.ToLower(n) == "hello" {
			commands++
			numberOfLines++
			fmt.Println("Hello to you too")
		} else if strings.ToLower(n) == "bye" {
			commands++
			numberOfLines++
			fmt.Println("Good bye to you as well!")
			break
		} else if strings.ToLower(n) == "q" {
			break
		} else {
			numberOfLines++
			fmt.Println("Unknows command! Try again!")
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Println("Commands entered:", commands)
	fmt.Println("Numbers of lines enteres:", numberOfLines)
}
