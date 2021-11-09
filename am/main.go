package am

import (
	"bufio"
	"fmt"
	"os"
)


func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SplitLines(s string) [][]byte {
	var count int

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && s[i+1] == 'n' {
			count++
		}
	}
	splitString := []byte(s)
	splitLines := make([][]byte, count+1)

	j := 0

	for i := 0; i < len(splitLines); i++ {
		for j < len(splitString) {
			if splitString[j] == 'n' && splitString[j-1] == '\\' {
				j++
				splitLines[i] = splitLines[i][:len(splitLines[i])-1]
				break
			}
			splitLines[i] = append(splitLines[i], splitString[j])
			j++
		}
	}
	return splitLines
}

func ColorRed(s map[int][]string) map[int][]string {
	colorRed := "\033[31m"

	fmt.Println(string(colorRed), s)

	return s
}

func ColorYellow(s map[int][]string) map[int][]string {
	colorYellow := "\033[33m"

	fmt.Println(string(colorYellow), s)

	return s
}

func ColorGreen(s map[int][]string) map[int][]string {
	colorGreen := "\033[32m"

	fmt.Println(string(colorGreen), s)

	return s
}

func ColorBlue(s map[int][]string) map[int][]string {
	colorBlue := "\033[34m"

	fmt.Println(string(colorBlue), s)

	return s
}

func ColorPurple(s map[int][]string) map[int][]string {
	colorPurple := "\033[35m"

	fmt.Println(string(colorPurple), s)

	return s
}

func ColorCyan(s map[int][]string) map[int][]string {
	colorCyan := "\033[36m"

	fmt.Println(string(colorCyan), s)

	return s
}

func ColorWhite(s map[int][]string) map[int][]string {
	colorWhite := "\033[37m"

	fmt.Println(string(colorWhite), s)

	return s
}

func ColorReset(s map[int][]string) map[int][]string {
	colorReset := "\033[0m"

	fmt.Println(string(colorReset), s)

	return s
}