package day20

import "testing"

var testInput = []string{
	"broadcaster -> a",
	"%a -> inv, con",
	"&inv -> b",
	"%b -> con",
	"&con -> output",
}

// button -low-> broadcaster
// broadcaster -low-> a
// a -high-> inv
// a -high-> con
// inv -low-> b
// con -high-> output
// b -high-> con
// con -low-> output

func TestMeasurePulses(t *testing.T) {
	expected := 32000000
	actual := MeasurePulses(testInput)

	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
