package am

import (
	"bufio"
	"fmt"
	"os"
)

// colorReset := "\033[0m"

// colorRed := "\033[31m"
// colorGreen := "\033[32m"
// colorYellow := "\033[33m"
// colorBlue := "\033[34m"
// colorPurple := "\033[35m"
// colorCyan := "\033[36m"
// colorWhite := "\033[37m"

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

func colorRed(s map[int][]string) map[int][]string {
	colorRed := "\033[31m"

	fmt.Println(string(colorRed), s)

	return s
}
