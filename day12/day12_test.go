package day12

import (
	"testing"
)

var testInput = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

// func TestSumOfArrangements(t *testing.T) {
// 	simple, complex := GetNumberOfArrangements(testInput[0])
// 	expectedResult := 1
// 	expectedComplexResult := 1

// 	if simple != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, simple)
// 	}

// 	if complex != expectedComplexResult {
// 		t.Errorf("Expected %+v, got %+v", expectedComplexResult, complex)
// 	}
// }

// func TestSumOfArrangements1(t *testing.T) {
// 	simple, complex := GetNumberOfArrangements(testInput[1])
// 	expectedResult := 4
// 	expectedComplexResult := 16384

// 	if simple != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, simple)
// 	}

// 	if complex != expectedComplexResult {
// 		t.Errorf("Expected %+v, got %+v", expectedComplexResult, complex)
// 	}
// }

// func TestSumOfArrangements2(t *testing.T) {
// 	simple, complex := GetNumberOfArrangements(testInput[2])
// 	expectedSimpleResult := 1
// 	expectedComplexResult := 1

// 	if simple != expectedSimpleResult {
// 		t.Errorf("Expected %+v, got %+v", expectedSimpleResult, simple)
// 	}

// 	if complex != expectedComplexResult {
// 		t.Errorf("Expected %+v, got %+v", expectedComplexResult, complex)
// 	}
// }

func TestSumOfArrangements3(t *testing.T) {
	simple, complex := GetNumberOfArrangements(testInput[3])
	expectedSimpleResult := 1
	expectedComplexResult := 16

	if simple != expectedSimpleResult {
		t.Errorf("Expected %+v, got %+v", expectedSimpleResult, simple)
	}

	if complex != expectedComplexResult {
		t.Errorf("Expected %+v, got %+v", expectedComplexResult, complex)
	}
}


// func TestTotalNumberOfArrangements(t *testing.T) {
// 	result := GetTotalNumberOfArrangements(testInput)
// 	expectedResult := 21

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// func TestTotalNumberOfArrangementsWithInput(t *testing.T) {
// 	godotenv.Load()
// 	expectedResult, _ := strconv.Atoi(os.Getenv("result_1"))

// 	Input, _ := utils.ReadInputFile("input.txt")
// 	result := GetTotalNumberOfArrangements(Input)

// 	if result != expectedResult {
// 		t.Errorf("Expected to be lower than %+v, got %+v", expectedResult, result)
// 	}
// }

// part 2

// func TestSumOfArrangementsWithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[0])
// 	expectedResult := 1

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// func TestSumOfArrangements1WithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[1], 4)
// 	expectedResult := 16384

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// func TestSumOfArrangements2WithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[2], 1)
// 	expectedResult := 1

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// func TestSumOfArrangements3WithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[3])
// 	expectedResult := 16

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// func TestSumOfArrangements4WithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[4], 5)
// 	expectedResult := 2500

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }
// func TestSumOfArrangements5WithCopies(t *testing.T) {
// 	result := GetNumberOfArrangements(testInput[5], 5)
// 	expectedResult := 506250

// 	if result != expectedResult {
// 		t.Errorf("Expected %+v, got %+v", expectedResult, result)
// 	}
// }

// 4 32 256 2048 16384
