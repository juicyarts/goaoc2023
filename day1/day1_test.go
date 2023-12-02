package day1

import (
	"bufio"
	"os"
	"reflect"
	"testing"
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

// Day 1 part 1

func TestSumCalibrationValue(t *testing.T) {
	got := FindFirstAndLastDigit("3abc2")
	want := 32

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValues1(t *testing.T) {
	got := SumCalbirationValues([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	want := 142

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValues2(t *testing.T) {
	got := SumCalbirationValues([]string{
		"1abc2",
		"1abc2",
	})

	want := 24

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValuesFromInput(t *testing.T) {
	Input, _ := ReadInputFile("input.txt")
	got := SumCalbirationValues(Input)
	want := 54597

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// Day 1 part 2

func TestFindAllIndexes(t *testing.T) {
	got := FindAllIndexesOfSubstringInString("two1two", "two")
	want := []int{0, 4}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValueWithWords(t *testing.T) {
	got := FindFirstAndLastDigitOrWord("fv9")
	want := 99

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValuesWithWords1(t *testing.T) {
	got := SumCalbirationValuesWithWords([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})

	want := 281

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumCalibrationValuesFromInputWithWords(t *testing.T) {
	Input, _ := ReadInputFile("input.txt")
	got := SumCalbirationValuesWithWords(Input)
	want := 54504

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
