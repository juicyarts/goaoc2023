package utils

import (
	"reflect"
	"testing"
)

func TestRotateAntiClockwise(t *testing.T) {
	localTestInput := []string{
		"ABCD",
		"EFGH",
	}

	expected := []string{
		"DH",
		"CG",
		"BF",
		"AE",
	}

	actual := RotateAntiClockWise(localTestInput)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRotateClockwise(t *testing.T) {
	localTestInput := []string{
		"ABCD",
		"EFGH",
	}

	expected := []string{
		"EA",
		"FB",
		"GC",
		"HD",
	}

	actual := RotateClockWise(localTestInput)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
