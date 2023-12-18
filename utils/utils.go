package utils

import (
	"bufio"
	"os"
)

func ReadInputFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func RotateAntiClockWise(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}
	result := make([]string, len(arr[0]))
	for i := 0; i < len(arr); i++ {
		for j := len(arr[i]) - 1; j >= 0; j-- {
			result[len(arr[i])-1-j] += string(arr[i][j])
		}
	}
	return result
}

func RotateClockWise(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}
	result := make([]string, len(arr[0]))
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			result[j] += string(arr[i][j])
		}
	}
	return result
}
