package day15

import (
	"aoc2023/utils"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var testInput = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7\n"

var testInputAsArray = []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}
var expectations = []int{
	30,
	253,
	97,
	47,
	14,
	180,
	9,
	197,
	48,
	214,
	231,
}

func TestHashing(t *testing.T) {
	for i, v := range testInputAsArray {
		result := Hash(v)
		if result != expectations[i] {
			t.Errorf("Hashing failed for %s", v)
			t.Errorf("Expected %+v got %v", expectations[i], result)
		}
	}
}
func TestHashing2(t *testing.T) {
	expectation := 52
	result := Hash("HASH")
	if result != expectation {
		t.Errorf("Hashing failed for 'HASH'")
		t.Errorf("Expected %+v got %v", expectation, result)
	}
}

func TestHashSum(t *testing.T) {
	expectation := 1320

	result := HashSum(testInput)
	if result != expectation {
		t.Errorf("Expected %+v got %v", expectation, result)
	}
}

func TestHashWithInput(t *testing.T) {
	godotenv.Load()
	expected, _ := strconv.Atoi(os.Getenv("result_1"))

	Input, _ := utils.ReadInputFile("input.txt")
	result := HashSum(Input[0])

	if result != expected {
		t.Errorf("Expected to be higher than %+v, got %+v", expected, result)
	}
}
