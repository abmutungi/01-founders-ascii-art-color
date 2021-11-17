package am

import (
	"bufio"
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

func TrimAtoi(s string) int {
	str := []rune(s)
	trim := 0

	for i := 0; i < len(str); i++ {
		// if !neg && trim == 0 && str[i] == '-' {
		// 	neg = true
		// }
		if str[i] >= '0' && str[i] <= '9' {
			trim *= 10
			trim += int(str[i] - 48)
		}
	}
	return trim
}
