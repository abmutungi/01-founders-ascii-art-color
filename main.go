package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"git.learn.01founders.co/abmutungi/ascii-art-color.git/am"
)

func main() {
	arg := os.Args

	// error handling
	if len(arg) != 3 {
		fmt.Print("Usage: go run . [STRING] [option]\n")
		fmt.Println()
		fmt.Println("EX: go run . something --color=<color>")
		os.Exit(0)
	}

	// defining argument [1] "input string" and [2] "font/banner"
	args := os.Args[1]
	args2 := os.Args[2]
	// args3, err := strconv.Atoi(os.Args[3])
	// _ = err

	// This tells it to print a new line if the arg is solely a new line.
	if args == "\\n" {
		fmt.Println()
	} else if args != "" {

		colourMap := map[string]string{
			"reset": "\033[0m",

			"red":    "\033[31m",
			"green":  "\033[32m",
			"yellow": "\033[33m",
			"blue":   "\033[34m",
			"purple": "\033[35m",
			"cyan":   "\033[36m",
			"white":  "\033[37m",
			"orange": "\033[38;5;214m",
		}
		// fmt.Println(colourMap)

		/*The else clause above tells the program to do nothing if
		the argument is an empty string with the rest of the program
		only running if the arg is not an empty string*/

		/* The func splitlines splits the string of the arg into
		a slice of slices split whenever there is a new line*/

		// splitLines := am.SplitLines(args)

		lines, err := am.ReadLines("standard.txt")
		if err != nil {
			log.Fatalf("ReadLines: %s", err)
		}

		/*The line below uses the make method to make a map
		and uses a start point of 32 to match up the ascii values
		of each character to the ascii version of the character*/
		charMap := make(map[int][]string)

		start := 32

		for i := 0; i < len(lines); i++ {
			// Tells it to add to start every 9 to match the chars
			if len(charMap[start]) == 9 {
				start++
			}
			charMap[start] = append(charMap[start], lines[i])
		}

		/*The j below refers to the index of each slice within string
		from the input (args) The i iterates up to 9
		to match the height of each character.*/

		// first condition allows user to add colour to the whole string
		if colourMap[args2[8:]] != "" && !strings.Contains(args2[8:], ":") {
			for i := 1; i < 9; i++ {
				for j := 0; j < len(args); j++ {
					// if j == args {
					fmt.Print(colourMap[args2[8:]], charMap[int(args[j])][i])
					// fmt.Print()
				}
				fmt.Println()
				fmt.Print(colourMap["reset"])

			} // second condition allows user to add colour to one character in string
			// user will need to use the index +1
		} else if colourMap[args2[8:(len(args2)-2)]] != "" && strings.Contains(args2[8:], ":") {
			for i := 1; i < 9; i++ {
				for j := 0; j < len(args); j++ {
					num := am.TrimAtoi(args2[8:])
					if num-1 == j {
						fmt.Print(colourMap[args2[8:(len(args2)-2)]], charMap[int(args[j])][i])
					} else {
						fmt.Print(colourMap["reset"], charMap[int(args[j])][i])
						// fmt.Print()
					}

				}
				fmt.Println()
				fmt.Print(colourMap["reset"])

			} // third condition allows user to add colour to a range of characters within string
			// user will need to use the index +1
		} else if colourMap[args2[8:(len(args2)-4)]] != "" && strings.Contains(args2[8:], "-") {

			num1 := args2[8 : len(args2)-1]
			num2 := args2[len(args2)-2:]
			first_num := am.TrimAtoi(num1)
			second_num := am.TrimAtoi(num2)

			for i := 1; i < 9; i++ {
				for j := 0; j < len(args); j++ {
					if j >= first_num-1 && j <= second_num-1 {
						fmt.Print(colourMap[args2[8:(len(args2)-4)]], charMap[int(args[j])][i])
					} else {
						fmt.Print(colourMap["reset"], charMap[int(args[j])][i])
					}
				}
				fmt.Println()
				fmt.Print(colourMap["reset"])
			}
		}
		// error handling
		if args2[8:] == "" {
			fmt.Print("Usage: go run . [STRING] [option]\n")
			fmt.Println()
			fmt.Println("EX: go run . something --color=<color>")
		}
	}
}
