package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.learn.01founders.co/abmutungi/ascii-art-color.git/am"
)

func main() {
	arg := os.Args

	// error handling

	if len(arg) != 4 {

		fmt.Print("Usage: go run . [STRING] [option]\n")
		fmt.Println()
		fmt.Println("EX: go run . something --color=<color>")
		os.Exit(0)
	}

	// defining argument [1] "input string" and [2] "font/banner"
	args := os.Args[1]
	args2 := os.Args[2]
	args3 := os.Args[3]
	numarg3, _ := strconv.Atoi(args3)
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
		}

		_ = colourMap

		// fmt.Println(colourMap)

		/*The else clause above tells the program to do nothing if
		the argument is an empty string with the rest of the program
		only running if the arg is not an empty string*/

		/* The func splitlines splits the string of the arg into
		a slice of slices split whenever there is a new line*/

		splitLines := am.SplitLines(args)

		// ColorMap := map[string]string{
		// 	"red":    "\033[31m",
		// 	"yellow": "\033[33m",
		// 	"green":  "\033[32m",
		// 	"blue":   "\033[34m",
		// 	"purple": "\033[35m",
		// 	"orange": "\033[16m",
		// 	"white":  "\033[37m",
		// 	"reset":  "\033[0m",
		// }

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

		// create empty string slice to append map to
		var eSlice []string

		/*The j below refers to the index of each slice within a
		created by splitlines, represented by val. The k represents
		the length of each individual slice. The i iterates up to 9
		to match the height of each character.*/

		for j, val := range splitLines {
			for i := 1; i < 9; i++ {
				for _, elem := range colourMap {
					for k := 0; k < len(val); k++ {
						if args2[8:] == colourMap[elem] && numarg3 == k {
							// if k == 2 {
							// fmt.Println("hello")
							eSlice = append((am.ColorBlue(eSlice)), charMap[int(splitLines[j][k])][i])
							// am.ColorReset(eSlice)
						} else {
							eSlice = append(eSlice, charMap[int(splitLines[j][k])][i])

							// fmt.Println(eSlice)
						}
						// eSlice = append(eSlice, charMap[int(splitLines[j][k])][i])

						// eSlice = append(eSlice, "\n")
						// fmt.Println()
					}
					// eSlice = append(eSlice, "\n")
				}
			}
			// am.ColorReset(eSlice)
			// }
			fSlice := strings.Join(eSlice, "")
			fmt.Println(fSlice)

			// if args2[8:] == "" {
			// 	fmt.Print("Usage: go run . [STRING] [option]\n")
			// 	fmt.Println()
			// 	fmt.Println("EX: go run . something --color=<color>")

			// }

			// if args2[8:] == "blue" {
			// 	fSlice = am.ColorBlue(fSlice)
			// }
			// if args2[8:] == "yellow" {
			// 	fSlice = am.ColorYellow(fSlice)
			// }

			// if args2[8:] == "white" {
			// 	fSlice = am.ColorWhite(fSlice)
			// }
			// // if args2[8:] == "orange" {
			// // 	fSlice = am.ColorOrange(fSlice)
			// // }

			// if args2[8:] == "reset" {
			// 	fSlice = am.ColorReset(fSlice)
			// }
			// if args2[8:] == "purple" {
			// 	fSlice = am.ColorPurple(fSlice)
			// }
		}
	}
}
