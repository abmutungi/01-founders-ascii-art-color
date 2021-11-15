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

// func ColorChar(s string) []byte {

// }

func ColorRed(s [] string) [] string {
	colorRed := "\033[31m"

	fmt.Println(colorRed, s)

	return s
}

func ColorOrange(s [] string) [] string {
	colorOrange := "\033[38;5;214m"

	fmt.Println(colorOrange, s)

	return s
}

func ColorYellow(s string) string {
	colorYellow := "\033[33m"

	fmt.Println(string(colorYellow), s)

	return s
}

func ColorGreen(s string) string {
	colorGreen := "\033[32m"

	fmt.Println(string(colorGreen), s)

	return s
}

func ColorBlue(s [] string) []string {
	colorBlue := "\033[34m"

	fmt.Println(string(colorBlue), s)

	return s
}

func ColorPurple(s string) string {
	colorPurple := "\033[35m"

	fmt.Println(string(colorPurple), s)

	return s
}

func ColorWhite(s string) string {
	colorWhite := "\033[37m"

	fmt.Println(string(colorWhite), s)

	return s
}

func ColorReset(s []string) []string {
	colorReset := "\033[0m"

	fmt.Println(colorReset, s)

	return s
}

func TrimAtoi(s string) int {
	neg := false
	str := []rune(s)
	trim := 0

	for i := 0; i < len(str); i++ {

		if !neg && trim == 0 && str[i] == '-' {
			neg = true
		}
		if str[i] >= '0' && str[i] <= '9' {
			trim *= 10
			trim += int(str[i] - 48)
		}
	}
	if neg {
		return trim * -1
	}
	return trim
}
